----------------------------- MODULE Chain ------------------------------

(***************************************************************************
 This module models the behavior of a chain running the IBC Core Client 
 Protocol.
****************************************************************************) 

EXTENDS Integers, FiniteSets, ICS02ClientHandlers, ICS02Definitions
        
CONSTANTS 
    MaxHeight, \* maximal chain height
    ChainID, \* chain identifier
    NrClients, \* number of clients that will be created on the chain
    ClientIDs \* a set of counterparty client IDs   

VARIABLES 
    chainStore, \* chain store, containing a client state for each client
    incomingDatagrams, \* set of incoming datagrams
    history \* history variable

vars == <<chainStore, incomingDatagrams, history>>
Heights == 1..MaxHeight \* set of possible heights of the chains in the system       

\* @type: (CHAINSTORE, Str) => Int;
GetClientNr(store, clientID) ==
    IF \E clientNr \in DOMAIN chainStore.clientStates :
            store.clientStates[clientNr].clientID = clientID
    THEN CHOOSE clientNr \in DOMAIN store.clientStates : 
            store.clientStates[clientNr].clientID = clientID 
    ELSE 0    

(***************************************************************************
 Client update operators
 ***************************************************************************)
\* Update the clients on chain with chainID, 
\* using the client datagrams generated by the relayer      
\* (Handler operators defined in ClientHandlers.tla)
LightClientUpdate(chainID, store, clientID, datagrams) == 
    \* create client 
    LET clientCreatedStore == HandleCreateClient(store, clientID, datagrams) IN
    \* update client
    LET clientUpdatedStore == HandleClientUpdate(clientCreatedStore, clientID, datagrams, MaxHeight) IN

    clientUpdatedStore

(***************************************************************************
 Chain actions
 ***************************************************************************)       
\* Advance the height of the chain until MaxHeight is reached
AdvanceChain ==
    /\ chainStore.height + 1 \in Heights
    /\ chainStore' = [chainStore EXCEPT !.height = chainStore.height + 1]
    /\ UNCHANGED <<incomingDatagrams, history>>

\* Handle the datagrams and update the chain state        
HandleIncomingDatagrams ==
    /\ incomingDatagrams /= {}
    /\ \E clientID \in ClientIDs : 
        /\ chainStore' = LightClientUpdate(ChainID, chainStore, clientID, incomingDatagrams)
        /\ history' = [history EXCEPT ![clientID] =
                            LET clientNr == GetClientNr(chainStore', clientID) IN 
                            IF /\ clientNr /= 0
                               /\ ~history[clientID].created
                               /\ chainStore.clientStates[clientNr].clientID = nullClientID
                               /\ chainStore'.clientStates[clientNr].clientID /= nullClientID
                            THEN [created |-> TRUE, updated |-> history[clientID].updated]
                            ELSE IF /\ clientNr /= 0
                                    /\ history[clientID].created
                                    /\ chainStore.clientStates[clientNr].heights /= chainStore'.clientStates[clientNr].heights
                                    /\ chainStore.clientStates[clientNr].heights \subseteq chainStore'.clientStates[clientNr].heights
                                THEN [created |-> history[clientID].created, updated |-> TRUE]
                                ELSE history[clientID]                               
                      ]
        /\ incomingDatagrams' = {dgr \in incomingDatagrams : dgr.clientID /= clientID}       

(***************************************************************************
 Specification
 ***************************************************************************)
\* Initial state predicate
\* Initially
\*  - each chain is initialized to InitChain (defined in RelayerDefinitions.tla)
\*  - pendingDatagrams for each chain is empty
\*  - the packetSeq is set to 1
Init == 
    /\ chainStore = ICS02InitChainStore(NrClients, ClientIDs)
    /\ incomingDatagrams = {}
    
\* Next state action
\* The chain either
\*  - advances its height
\*  - receives datagrams and updates its state
\*  - sends a packet if the appPacketSeq is not bigger than MaxPacketSeq
\*  - acknowledges a packet
Next ==
    \/ AdvanceChain 
    \/ HandleIncomingDatagrams
    \/ UNCHANGED vars
        
Fairness ==
    /\ WF_vars(AdvanceChain)
    /\ WF_vars(HandleIncomingDatagrams)        
        
(***************************************************************************
 Invariants
 ***************************************************************************)
\* Type invariant   
\* ChainStores and Datagrams are defined in RelayerDefinitions.tla        
TypeOK ==    
    /\ chainStore \in ChainStores(NrClients, ClientIDs, MaxHeight)
    /\ incomingDatagrams \in SUBSET Datagrams(ClientIDs, MaxHeight)
    
\* two clients with the same ID cannot be created    
CreatedClientsHaveDifferentIDs ==
    (\A clientNr \in 1..NrClients : 
        chainStore.clientStates[clientNr].clientID /= nullClientID)
        => (\A clientNr1 \in 1..NrClients : \A clientNr2 \in 1..NrClients :
                clientNr1 /= clientNr2 
                    => chainStore.clientStates[clientNr1].clientID /= 
                       chainStore.clientStates[clientNr2].clientID)
        
\* only created clients can be updated
UpdatedClientsAreCreated ==
    \A clID \in ClientIDs : 
        history[clID].updated => history[clID].created

(***************************************************************************
 Properties
 ***************************************************************************)    
\* it ALWAYS holds that the height of the chain does not EVENTUALLY decrease
HeightDoesntDecrease ==
    [](\A h \in Heights : chainStore.height = h 
        => <>(chainStore.height >= h))
        
=============================================================================
\* Modification History
\* Last modified Thu Apr 15 12:17:59 CEST 2021 by ilinastoilkovska
\* Created Fri Jun 05 16:56:21 CET 2020 by ilinastoilkovska
