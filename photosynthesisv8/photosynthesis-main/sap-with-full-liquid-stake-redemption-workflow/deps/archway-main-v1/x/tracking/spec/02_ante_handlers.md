<!--
order: 2
-->

# AnteHandlers

Section describes the module ante handlers.

## TxGasTrackingDecorator

The [TxGasTrackingDecorator](../ante/tracking.go#L15) handler kickstarts a
transaction tracking by creating an empty [TxInfo](01\_state.md#TxInfo).
