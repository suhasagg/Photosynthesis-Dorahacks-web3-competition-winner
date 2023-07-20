(module
  (type (;0;) (func (param i32 i32)))
  (type (;1;) (func (param i32 i32) (result i32)))
  (type (;2;) (func (param i32) (result i32)))
  (type (;3;) (func (param i32 i32 i32) (result i32)))
  (type (;4;) (func (param i32)))
  (type (;5;) (func (param i32) (result i64)))
  (type (;6;) (func (param i32 i32 i32)))
  (type (;7;) (func (result i32)))
  (type (;8;) (func))
  (type (;9;) (func (param i32 i32 i32 i32)))
  (type (;10;) (func (param i32 i32 i32 i32 i32)))
  (type (;11;) (func (param i32 i32 i32 i32) (result i32)))
  (func (;0;) (type 7) (result i32)
    i32.const 1)
  (func (;1;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 13)
  (func (;2;) (type 5) (param i32) (result i64)
    i64.const -4001154616061892501)
  (func (;3;) (type 5) (param i32) (result i64)
    i64.const -8527728395957036344)
  (func (;4;) (type 5) (param i32) (result i64)
    i64.const -7968955924275895985)
  (func (;5;) (type 6) (param i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 3
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        local.get 1
        local.get 2
        i32.add
        local.tee 1
        i32.gt_u
        br_if 0 (;@2;)
        local.get 0
        i32.load
        local.tee 2
        i32.const 1
        i32.shl
        local.tee 4
        local.get 1
        local.get 1
        local.get 4
        i32.lt_u
        select
        local.tee 1
        i32.const 8
        local.get 1
        i32.const 8
        i32.gt_u
        select
        local.tee 1
        i32.const -1
        i32.xor
        i32.const 31
        i32.shr_u
        local.set 4
        block  ;; label = @3
          local.get 2
          if  ;; label = @4
            local.get 3
            i32.const 1
            i32.store offset=24
            local.get 3
            local.get 2
            i32.store offset=20
            local.get 3
            local.get 0
            i32.const 4
            i32.add
            i32.load
            i32.store offset=16
            br 1 (;@3;)
          end
          local.get 3
          i32.const 0
          i32.store offset=24
        end
        local.get 3
        local.get 1
        local.get 4
        local.get 3
        i32.const 16
        i32.add
        call 12
        local.get 3
        i32.load offset=4
        local.set 2
        local.get 3
        i32.load
        i32.eqz
        if  ;; label = @3
          local.get 0
          local.get 1
          i32.store
          local.get 0
          local.get 2
          i32.store offset=4
          br 2 (;@1;)
        end
        local.get 3
        i32.const 8
        i32.add
        i32.load
        local.tee 0
        i32.const -2147483647
        i32.eq
        br_if 1 (;@1;)
        local.get 0
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        local.get 0
        call 47
        unreachable
      end
      call 48
      unreachable
    end
    local.get 3
    i32.const 32
    i32.add
    global.set 0)
  (func (;6;) (type 4) (param i32)
    nop)
  (func (;7;) (type 4) (param i32)
    local.get 0
    i32.load
    if  ;; label = @1
      local.get 0
      i32.const 4
      i32.add
      i32.load
      call 19
    end)
  (func (;8;) (type 4) (param i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.const 4
      i32.add
      i32.load
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.load
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      call 19
    end)
  (func (;9;) (type 1) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 0
    i32.load
    local.set 0
    block  ;; label = @1
      block (result i32)  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.const 128
          i32.ge_u
          if  ;; label = @4
            local.get 3
            i32.const 0
            i32.store offset=12
            local.get 1
            i32.const 2048
            i32.ge_u
            br_if 1 (;@3;)
            local.get 3
            local.get 1
            i32.const 63
            i32.and
            i32.const 128
            i32.or
            i32.store8 offset=13
            local.get 3
            local.get 1
            i32.const 6
            i32.shr_u
            i32.const 192
            i32.or
            i32.store8 offset=12
            i32.const 2
            br 2 (;@2;)
          end
          local.get 0
          i32.load offset=8
          local.tee 2
          local.get 0
          i32.load
          i32.eq
          if  ;; label = @4
            global.get 0
            i32.const 32
            i32.sub
            local.tee 4
            global.set 0
            block  ;; label = @5
              block  ;; label = @6
                local.get 2
                i32.const 1
                i32.add
                local.tee 2
                i32.eqz
                br_if 0 (;@6;)
                local.get 0
                i32.load
                local.tee 5
                i32.const 1
                i32.shl
                local.tee 6
                local.get 2
                local.get 2
                local.get 6
                i32.lt_u
                select
                local.tee 2
                i32.const 8
                local.get 2
                i32.const 8
                i32.gt_u
                select
                local.tee 2
                i32.const -1
                i32.xor
                i32.const 31
                i32.shr_u
                local.set 6
                block  ;; label = @7
                  local.get 5
                  if  ;; label = @8
                    local.get 4
                    i32.const 1
                    i32.store offset=24
                    local.get 4
                    local.get 5
                    i32.store offset=20
                    local.get 4
                    local.get 0
                    i32.const 4
                    i32.add
                    i32.load
                    i32.store offset=16
                    br 1 (;@7;)
                  end
                  local.get 4
                  i32.const 0
                  i32.store offset=24
                end
                local.get 4
                local.get 2
                local.get 6
                local.get 4
                i32.const 16
                i32.add
                call 12
                local.get 4
                i32.load offset=4
                local.set 5
                local.get 4
                i32.load
                i32.eqz
                if  ;; label = @7
                  local.get 0
                  local.get 2
                  i32.store
                  local.get 0
                  local.get 5
                  i32.store offset=4
                  br 2 (;@5;)
                end
                local.get 4
                i32.const 8
                i32.add
                i32.load
                local.tee 2
                i32.const -2147483647
                i32.eq
                br_if 1 (;@5;)
                local.get 2
                i32.eqz
                br_if 0 (;@6;)
                local.get 5
                local.get 2
                call 47
                unreachable
              end
              call 48
              unreachable
            end
            local.get 4
            i32.const 32
            i32.add
            global.set 0
            local.get 0
            i32.load offset=8
            local.set 2
          end
          local.get 0
          local.get 2
          i32.const 1
          i32.add
          i32.store offset=8
          local.get 0
          i32.load offset=4
          local.get 2
          i32.add
          local.get 1
          i32.store8
          br 2 (;@1;)
        end
        local.get 1
        i32.const 65536
        i32.ge_u
        if  ;; label = @3
          local.get 3
          local.get 1
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=15
          local.get 3
          local.get 1
          i32.const 6
          i32.shr_u
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=14
          local.get 3
          local.get 1
          i32.const 12
          i32.shr_u
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=13
          local.get 3
          local.get 1
          i32.const 18
          i32.shr_u
          i32.const 7
          i32.and
          i32.const 240
          i32.or
          i32.store8 offset=12
          i32.const 4
          br 1 (;@2;)
        end
        local.get 3
        local.get 1
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=14
        local.get 3
        local.get 1
        i32.const 12
        i32.shr_u
        i32.const 224
        i32.or
        i32.store8 offset=12
        local.get 3
        local.get 1
        i32.const 6
        i32.shr_u
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=13
        i32.const 3
      end
      local.set 1
      local.get 1
      local.get 0
      i32.load
      local.get 0
      i32.load offset=8
      local.tee 2
      i32.sub
      i32.gt_u
      if  ;; label = @2
        local.get 0
        local.get 2
        local.get 1
        call 5
        local.get 0
        i32.load offset=8
        local.set 2
      end
      local.get 0
      i32.load offset=4
      local.get 2
      i32.add
      local.get 3
      i32.const 12
      i32.add
      local.get 1
      call 55
      drop
      local.get 0
      local.get 1
      local.get 2
      i32.add
      i32.store offset=8
    end
    local.get 3
    i32.const 16
    i32.add
    global.set 0
    i32.const 0)
  (func (;10;) (type 1) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.load
    i32.store offset=4
    local.get 2
    i32.const 24
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 16
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 1
    i64.load align=4
    i64.store offset=8
    local.get 2
    i32.const 4
    i32.add
    local.get 2
    i32.const 8
    i32.add
    call 53
    local.get 2
    i32.const 32
    i32.add
    global.set 0)
  (func (;11;) (type 3) (param i32 i32 i32) (result i32)
    (local i32)
    local.get 2
    local.get 0
    i32.load
    local.tee 0
    i32.load
    local.get 0
    i32.load offset=8
    local.tee 3
    i32.sub
    i32.gt_u
    if  ;; label = @1
      local.get 0
      local.get 3
      local.get 2
      call 5
      local.get 0
      i32.load offset=8
      local.set 3
    end
    local.get 0
    i32.load offset=4
    local.get 3
    i32.add
    local.get 1
    local.get 2
    call 55
    drop
    local.get 0
    local.get 2
    local.get 3
    i32.add
    i32.store offset=8
    i32.const 0)
  (func (;12;) (type 9) (param i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      local.get 2
      if  ;; label = @2
        block (result i32)  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 1
                i32.const 0
                i32.ge_s
                if  ;; label = @7
                  local.get 3
                  i32.load offset=8
                  i32.eqz
                  br_if 2 (;@5;)
                  local.get 3
                  i32.load offset=4
                  local.tee 4
                  br_if 1 (;@6;)
                  local.get 1
                  br_if 3 (;@4;)
                  local.get 2
                  br 4 (;@3;)
                end
                local.get 0
                i32.const 8
                i32.add
                i32.const 0
                i32.store
                br 5 (;@1;)
              end
              local.get 3
              i32.load
              local.set 9
              block (result i32)  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 2
                      i32.const 9
                      i32.ge_u
                      if  ;; label = @10
                        local.get 1
                        local.get 2
                        call 13
                        local.tee 11
                        br_if 1 (;@9;)
                        i32.const 0
                        br 4 (;@6;)
                      end
                      i32.const 8
                      i32.const 8
                      call 26
                      local.set 7
                      i32.const 20
                      i32.const 8
                      call 26
                      local.set 6
                      i32.const 16
                      i32.const 8
                      call 26
                      local.set 3
                      i32.const 0
                      i32.const 16
                      i32.const 8
                      call 26
                      i32.const 2
                      i32.shl
                      i32.sub
                      local.tee 4
                      i32.const -65536
                      local.get 3
                      local.get 6
                      local.get 7
                      i32.add
                      i32.add
                      i32.sub
                      i32.const -9
                      i32.and
                      i32.const 3
                      i32.sub
                      local.tee 3
                      local.get 3
                      local.get 4
                      i32.gt_u
                      select
                      local.get 1
                      i32.le_u
                      br_if 1 (;@8;)
                      i32.const 16
                      local.get 1
                      i32.const 4
                      i32.add
                      i32.const 16
                      i32.const 8
                      call 26
                      i32.const 5
                      i32.sub
                      local.get 1
                      i32.gt_u
                      select
                      i32.const 8
                      call 26
                      local.set 5
                      local.get 9
                      call 42
                      local.tee 4
                      local.get 4
                      call 30
                      local.tee 3
                      call 39
                      local.set 8
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    local.get 4
                                    call 33
                                    i32.eqz
                                    if  ;; label = @17
                                      local.get 3
                                      local.get 5
                                      i32.ge_u
                                      br_if 1 (;@16;)
                                      local.get 8
                                      i32.const 1049844
                                      i32.load
                                      i32.eq
                                      br_if 2 (;@15;)
                                      local.get 8
                                      i32.const 1049840
                                      i32.load
                                      i32.eq
                                      br_if 3 (;@14;)
                                      local.get 8
                                      call 31
                                      br_if 7 (;@10;)
                                      local.get 8
                                      call 30
                                      local.tee 10
                                      local.get 3
                                      i32.add
                                      local.tee 7
                                      local.get 5
                                      i32.lt_u
                                      br_if 7 (;@10;)
                                      local.get 7
                                      local.get 5
                                      i32.sub
                                      local.set 12
                                      local.get 10
                                      i32.const 256
                                      i32.lt_u
                                      br_if 4 (;@13;)
                                      local.get 8
                                      call 16
                                      br 5 (;@12;)
                                    end
                                    local.get 4
                                    call 30
                                    local.set 3
                                    local.get 5
                                    i32.const 256
                                    i32.lt_u
                                    br_if 6 (;@10;)
                                    local.get 3
                                    local.get 5
                                    i32.sub
                                    i32.const 131073
                                    i32.lt_u
                                    local.get 5
                                    i32.const 4
                                    i32.add
                                    local.get 3
                                    i32.le_u
                                    i32.and
                                    br_if 5 (;@11;)
                                    local.get 3
                                    local.get 4
                                    i32.load
                                    local.tee 3
                                    i32.add
                                    i32.const 16
                                    i32.add
                                    local.set 7
                                    local.get 5
                                    i32.const 31
                                    i32.add
                                    i32.const 65536
                                    call 26
                                    local.set 10
                                    br 6 (;@10;)
                                  end
                                  i32.const 16
                                  i32.const 8
                                  call 26
                                  local.get 3
                                  local.get 5
                                  i32.sub
                                  local.tee 6
                                  i32.gt_u
                                  br_if 4 (;@11;)
                                  local.get 4
                                  local.get 5
                                  call 39
                                  local.set 3
                                  local.get 4
                                  local.get 5
                                  call 34
                                  local.get 3
                                  local.get 6
                                  call 34
                                  local.get 3
                                  local.get 6
                                  call 15
                                  br 4 (;@11;)
                                end
                                i32.const 1049836
                                i32.load
                                local.get 3
                                i32.add
                                local.tee 3
                                local.get 5
                                i32.le_u
                                br_if 4 (;@10;)
                                local.get 4
                                local.get 5
                                call 39
                                local.set 6
                                local.get 4
                                local.get 5
                                call 34
                                local.get 6
                                local.get 3
                                local.get 5
                                i32.sub
                                local.tee 3
                                i32.const 1
                                i32.or
                                i32.store offset=4
                                i32.const 1049836
                                local.get 3
                                i32.store
                                i32.const 1049844
                                local.get 6
                                i32.store
                                br 3 (;@11;)
                              end
                              i32.const 1049832
                              i32.load
                              local.get 3
                              i32.add
                              local.tee 3
                              local.get 5
                              i32.lt_u
                              br_if 3 (;@10;)
                              block  ;; label = @14
                                i32.const 16
                                i32.const 8
                                call 26
                                local.get 3
                                local.get 5
                                i32.sub
                                local.tee 7
                                i32.gt_u
                                if  ;; label = @15
                                  local.get 4
                                  local.get 3
                                  call 34
                                  i32.const 0
                                  local.set 7
                                  i32.const 0
                                  local.set 6
                                  br 1 (;@14;)
                                end
                                local.get 4
                                local.get 5
                                call 39
                                local.tee 6
                                local.get 7
                                call 39
                                local.set 3
                                local.get 4
                                local.get 5
                                call 34
                                local.get 6
                                local.get 7
                                call 37
                                local.get 3
                                local.get 3
                                i32.load offset=4
                                i32.const -2
                                i32.and
                                i32.store offset=4
                              end
                              i32.const 1049840
                              local.get 6
                              i32.store
                              i32.const 1049832
                              local.get 7
                              i32.store
                              br 2 (;@11;)
                            end
                            local.get 8
                            i32.const 12
                            i32.add
                            i32.load
                            local.tee 6
                            local.get 8
                            i32.const 8
                            i32.add
                            i32.load
                            local.tee 3
                            i32.ne
                            if  ;; label = @13
                              local.get 3
                              local.get 6
                              i32.store offset=12
                              local.get 6
                              local.get 3
                              i32.store offset=8
                              br 1 (;@12;)
                            end
                            i32.const 1049824
                            i32.const 1049824
                            i32.load
                            i32.const -2
                            local.get 10
                            i32.const 3
                            i32.shr_u
                            i32.rotl
                            i32.and
                            i32.store
                          end
                          i32.const 16
                          i32.const 8
                          call 26
                          local.get 12
                          i32.le_u
                          if  ;; label = @12
                            local.get 4
                            local.get 5
                            call 39
                            local.set 3
                            local.get 4
                            local.get 5
                            call 34
                            local.get 3
                            local.get 12
                            call 34
                            local.get 3
                            local.get 12
                            call 15
                            br 1 (;@11;)
                          end
                          local.get 4
                          local.get 7
                          call 34
                        end
                        local.get 4
                        br_if 3 (;@7;)
                      end
                      local.get 1
                      call 14
                      local.tee 3
                      i32.eqz
                      br_if 1 (;@8;)
                      local.get 3
                      local.get 9
                      local.get 4
                      call 30
                      i32.const -8
                      i32.const -4
                      local.get 4
                      call 33
                      select
                      i32.add
                      local.tee 3
                      local.get 1
                      local.get 1
                      local.get 3
                      i32.gt_u
                      select
                      call 55
                      local.get 9
                      call 19
                      br 3 (;@6;)
                    end
                    local.get 11
                    local.get 9
                    local.get 4
                    local.get 1
                    local.get 1
                    local.get 4
                    i32.gt_u
                    select
                    call 55
                    drop
                    local.get 9
                    call 19
                  end
                  local.get 11
                  br 1 (;@6;)
                end
                local.get 4
                call 33
                drop
                local.get 4
                call 41
              end
              br 2 (;@3;)
            end
            local.get 1
            br_if 0 (;@4;)
            local.get 2
            br 1 (;@3;)
          end
          local.get 1
          local.get 2
          call 1
        end
        local.tee 3
        if  ;; label = @3
          local.get 0
          local.get 3
          i32.store offset=4
          local.get 0
          i32.const 8
          i32.add
          local.get 1
          i32.store
          local.get 0
          i32.const 0
          i32.store
          return
        end
        local.get 0
        local.get 1
        i32.store offset=4
        local.get 0
        i32.const 8
        i32.add
        local.get 2
        i32.store
        br 1 (;@1;)
      end
      local.get 0
      local.get 1
      i32.store offset=4
      local.get 0
      i32.const 8
      i32.add
      i32.const 0
      i32.store
    end
    local.get 0
    i32.const 1
    i32.store)
  (func (;13;) (type 1) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.const 9
            i32.ge_u
            if  ;; label = @5
              i32.const 16
              i32.const 8
              call 26
              local.get 1
              i32.gt_u
              br_if 1 (;@4;)
              br 2 (;@3;)
            end
            local.get 0
            call 14
            local.set 4
            br 2 (;@2;)
          end
          i32.const 16
          i32.const 8
          call 26
          local.set 1
        end
        i32.const 8
        i32.const 8
        call 26
        local.set 3
        i32.const 20
        i32.const 8
        call 26
        local.set 2
        i32.const 16
        i32.const 8
        call 26
        local.set 5
        i32.const 0
        i32.const 16
        i32.const 8
        call 26
        i32.const 2
        i32.shl
        i32.sub
        local.tee 6
        i32.const -65536
        local.get 5
        local.get 2
        local.get 3
        i32.add
        i32.add
        i32.sub
        i32.const -9
        i32.and
        i32.const 3
        i32.sub
        local.tee 3
        local.get 3
        local.get 6
        i32.gt_u
        select
        local.get 1
        i32.sub
        local.get 0
        i32.le_u
        br_if 0 (;@2;)
        local.get 1
        i32.const 16
        local.get 0
        i32.const 4
        i32.add
        i32.const 16
        i32.const 8
        call 26
        i32.const 5
        i32.sub
        local.get 0
        i32.gt_u
        select
        i32.const 8
        call 26
        local.tee 3
        i32.add
        i32.const 16
        i32.const 8
        call 26
        i32.add
        i32.const 4
        i32.sub
        call 14
        local.tee 2
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        call 42
        local.set 0
        block  ;; label = @3
          local.get 1
          i32.const 1
          i32.sub
          local.tee 4
          local.get 2
          i32.and
          i32.eqz
          if  ;; label = @4
            local.get 0
            local.set 1
            br 1 (;@3;)
          end
          local.get 2
          local.get 4
          i32.add
          i32.const 0
          local.get 1
          i32.sub
          i32.and
          call 42
          local.set 2
          i32.const 16
          i32.const 8
          call 26
          local.set 4
          local.get 0
          call 30
          local.get 2
          i32.const 0
          local.get 1
          local.get 2
          local.get 0
          i32.sub
          local.get 4
          i32.gt_u
          select
          i32.add
          local.tee 1
          local.get 0
          i32.sub
          local.tee 2
          i32.sub
          local.set 4
          local.get 0
          call 33
          i32.eqz
          if  ;; label = @4
            local.get 1
            local.get 4
            call 34
            local.get 0
            local.get 2
            call 34
            local.get 0
            local.get 2
            call 15
            br 1 (;@3;)
          end
          local.get 0
          i32.load
          local.set 0
          local.get 1
          local.get 4
          i32.store offset=4
          local.get 1
          local.get 0
          local.get 2
          i32.add
          i32.store
        end
        local.get 1
        call 33
        br_if 1 (;@1;)
        local.get 1
        call 30
        local.tee 2
        i32.const 16
        i32.const 8
        call 26
        local.get 3
        i32.add
        i32.le_u
        br_if 1 (;@1;)
        local.get 1
        local.get 3
        call 39
        local.set 0
        local.get 1
        local.get 3
        call 34
        local.get 0
        local.get 2
        local.get 3
        i32.sub
        local.tee 3
        call 34
        local.get 0
        local.get 3
        call 15
        br 1 (;@1;)
      end
      local.get 4
      return
    end
    local.get 1
    call 41
    local.get 1
    call 33
    drop)
  (func (;14;) (type 2) (param i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 8
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.const 245
        i32.ge_u
        if  ;; label = @3
          i32.const 8
          i32.const 8
          call 26
          local.set 2
          i32.const 20
          i32.const 8
          call 26
          local.set 3
          i32.const 16
          i32.const 8
          call 26
          local.set 5
          i32.const 0
          i32.const 16
          i32.const 8
          call 26
          i32.const 2
          i32.shl
          i32.sub
          local.tee 4
          i32.const -65536
          local.get 5
          local.get 2
          local.get 3
          i32.add
          i32.add
          i32.sub
          i32.const -9
          i32.and
          i32.const 3
          i32.sub
          local.tee 2
          local.get 2
          local.get 4
          i32.gt_u
          select
          local.get 0
          i32.le_u
          br_if 2 (;@1;)
          local.get 0
          i32.const 4
          i32.add
          i32.const 8
          call 26
          local.set 4
          i32.const 1049828
          i32.load
          i32.eqz
          br_if 1 (;@2;)
          i32.const 0
          local.get 4
          i32.sub
          local.set 1
          block  ;; label = @4
            block  ;; label = @5
              block (result i32)  ;; label = @6
                i32.const 0
                local.get 4
                i32.const 256
                i32.lt_u
                br_if 0 (;@6;)
                drop
                i32.const 31
                local.get 4
                i32.const 16777215
                i32.gt_u
                br_if 0 (;@6;)
                drop
                local.get 4
                i32.const 6
                local.get 4
                i32.const 8
                i32.shr_u
                i32.clz
                local.tee 0
                i32.sub
                i32.shr_u
                i32.const 1
                i32.and
                local.get 0
                i32.const 1
                i32.shl
                i32.sub
                i32.const 62
                i32.add
              end
              local.tee 7
              i32.const 2
              i32.shl
              i32.const 1049416
              i32.add
              i32.load
              local.tee 0
              if  ;; label = @6
                local.get 4
                local.get 7
                call 29
                i32.shl
                local.set 6
                i32.const 0
                local.set 3
                i32.const 0
                local.set 2
                loop  ;; label = @7
                  block  ;; label = @8
                    local.get 0
                    call 30
                    local.tee 5
                    local.get 4
                    i32.lt_u
                    br_if 0 (;@8;)
                    local.get 5
                    local.get 4
                    i32.sub
                    local.tee 5
                    local.get 1
                    i32.ge_u
                    br_if 0 (;@8;)
                    local.get 0
                    local.set 2
                    local.get 5
                    local.tee 1
                    br_if 0 (;@8;)
                    i32.const 0
                    local.set 1
                    br 3 (;@5;)
                  end
                  local.get 0
                  i32.const 20
                  i32.add
                  i32.load
                  local.tee 5
                  local.get 3
                  local.get 5
                  local.get 0
                  local.get 6
                  i32.const 29
                  i32.shr_u
                  i32.const 4
                  i32.and
                  i32.add
                  i32.const 16
                  i32.add
                  i32.load
                  local.tee 0
                  i32.ne
                  select
                  local.get 3
                  local.get 5
                  select
                  local.set 3
                  local.get 6
                  i32.const 1
                  i32.shl
                  local.set 6
                  local.get 0
                  br_if 0 (;@7;)
                end
                local.get 3
                if  ;; label = @7
                  local.get 3
                  local.set 0
                  br 2 (;@5;)
                end
                local.get 2
                br_if 2 (;@4;)
              end
              i32.const 0
              local.set 2
              i32.const 1
              local.get 7
              i32.shl
              call 27
              i32.const 1049828
              i32.load
              i32.and
              local.tee 0
              i32.eqz
              br_if 3 (;@2;)
              local.get 0
              call 28
              i32.ctz
              i32.const 2
              i32.shl
              i32.const 1049416
              i32.add
              i32.load
              local.tee 0
              i32.eqz
              br_if 3 (;@2;)
            end
            loop  ;; label = @5
              local.get 0
              local.get 2
              local.get 0
              call 30
              local.tee 2
              local.get 4
              i32.ge_u
              local.get 2
              local.get 4
              i32.sub
              local.tee 3
              local.get 1
              i32.lt_u
              i32.and
              local.tee 5
              select
              local.set 2
              local.get 3
              local.get 1
              local.get 5
              select
              local.set 1
              local.get 0
              call 43
              local.tee 0
              br_if 0 (;@5;)
            end
            local.get 2
            i32.eqz
            br_if 2 (;@2;)
          end
          local.get 4
          i32.const 1049832
          i32.load
          local.tee 0
          i32.le_u
          local.get 1
          local.get 0
          local.get 4
          i32.sub
          i32.ge_u
          i32.and
          br_if 1 (;@2;)
          local.get 2
          local.get 4
          call 39
          local.set 0
          local.get 2
          call 16
          block  ;; label = @4
            i32.const 16
            i32.const 8
            call 26
            local.get 1
            i32.le_u
            if  ;; label = @5
              local.get 2
              local.get 4
              call 36
              local.get 0
              local.get 1
              call 37
              local.get 1
              i32.const 256
              i32.ge_u
              if  ;; label = @6
                local.get 0
                local.get 1
                call 17
                br 2 (;@4;)
              end
              local.get 1
              i32.const -8
              i32.and
              i32.const 1049560
              i32.add
              local.set 3
              block (result i32)  ;; label = @6
                i32.const 1049824
                i32.load
                local.tee 5
                i32.const 1
                local.get 1
                i32.const 3
                i32.shr_u
                i32.shl
                local.tee 1
                i32.and
                if  ;; label = @7
                  local.get 3
                  i32.load offset=8
                  br 1 (;@6;)
                end
                i32.const 1049824
                local.get 1
                local.get 5
                i32.or
                i32.store
                local.get 3
              end
              local.set 1
              local.get 3
              local.get 0
              i32.store offset=8
              local.get 1
              local.get 0
              i32.store offset=12
              local.get 0
              local.get 3
              i32.store offset=12
              local.get 0
              local.get 1
              i32.store offset=8
              br 1 (;@4;)
            end
            local.get 2
            local.get 1
            local.get 4
            i32.add
            call 35
          end
          local.get 2
          call 41
          local.tee 1
          i32.eqz
          br_if 1 (;@2;)
          br 2 (;@1;)
        end
        i32.const 16
        local.get 0
        i32.const 4
        i32.add
        i32.const 16
        i32.const 8
        call 26
        i32.const 5
        i32.sub
        local.get 0
        i32.gt_u
        select
        i32.const 8
        call 26
        local.set 4
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block (result i32)  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    i32.const 1049824
                    i32.load
                    local.tee 5
                    local.get 4
                    i32.const 3
                    i32.shr_u
                    local.tee 1
                    i32.shr_u
                    local.tee 0
                    i32.const 3
                    i32.and
                    i32.eqz
                    if  ;; label = @9
                      local.get 4
                      i32.const 1049832
                      i32.load
                      i32.le_u
                      br_if 7 (;@2;)
                      local.get 0
                      br_if 1 (;@8;)
                      i32.const 1049828
                      i32.load
                      local.tee 0
                      i32.eqz
                      br_if 7 (;@2;)
                      local.get 0
                      call 28
                      i32.ctz
                      i32.const 2
                      i32.shl
                      i32.const 1049416
                      i32.add
                      i32.load
                      local.tee 2
                      call 30
                      local.get 4
                      i32.sub
                      local.set 1
                      local.get 2
                      call 43
                      local.tee 0
                      if  ;; label = @10
                        loop  ;; label = @11
                          local.get 0
                          call 30
                          local.get 4
                          i32.sub
                          local.tee 3
                          local.get 1
                          local.get 1
                          local.get 3
                          i32.gt_u
                          local.tee 3
                          select
                          local.set 1
                          local.get 0
                          local.get 2
                          local.get 3
                          select
                          local.set 2
                          local.get 0
                          call 43
                          local.tee 0
                          br_if 0 (;@11;)
                        end
                      end
                      local.get 2
                      local.get 4
                      call 39
                      local.set 5
                      local.get 2
                      call 16
                      i32.const 16
                      i32.const 8
                      call 26
                      local.get 1
                      i32.gt_u
                      br_if 5 (;@4;)
                      local.get 2
                      local.get 4
                      call 36
                      local.get 5
                      local.get 1
                      call 37
                      i32.const 1049832
                      i32.load
                      local.tee 6
                      i32.eqz
                      br_if 4 (;@5;)
                      local.get 6
                      i32.const -8
                      i32.and
                      i32.const 1049560
                      i32.add
                      local.set 0
                      i32.const 1049840
                      i32.load
                      local.set 3
                      i32.const 1049824
                      i32.load
                      local.tee 7
                      i32.const 1
                      local.get 6
                      i32.const 3
                      i32.shr_u
                      i32.shl
                      local.tee 6
                      i32.and
                      i32.eqz
                      br_if 2 (;@7;)
                      local.get 0
                      i32.load offset=8
                      br 3 (;@6;)
                    end
                    block  ;; label = @9
                      local.get 0
                      i32.const -1
                      i32.xor
                      i32.const 1
                      i32.and
                      local.get 1
                      i32.add
                      local.tee 0
                      i32.const 3
                      i32.shl
                      local.tee 3
                      i32.const 1049568
                      i32.add
                      i32.load
                      local.tee 1
                      i32.const 8
                      i32.add
                      i32.load
                      local.tee 2
                      local.get 3
                      i32.const 1049560
                      i32.add
                      local.tee 3
                      i32.ne
                      if  ;; label = @10
                        local.get 2
                        local.get 3
                        i32.store offset=12
                        local.get 3
                        local.get 2
                        i32.store offset=8
                        br 1 (;@9;)
                      end
                      i32.const 1049824
                      local.get 5
                      i32.const -2
                      local.get 0
                      i32.rotl
                      i32.and
                      i32.store
                    end
                    local.get 1
                    local.get 0
                    i32.const 3
                    i32.shl
                    call 35
                    local.get 1
                    call 41
                    local.set 1
                    br 7 (;@1;)
                  end
                  block  ;; label = @8
                    i32.const 1
                    local.get 1
                    i32.const 31
                    i32.and
                    local.tee 1
                    i32.shl
                    call 27
                    local.get 0
                    local.get 1
                    i32.shl
                    i32.and
                    call 28
                    i32.ctz
                    local.tee 0
                    i32.const 3
                    i32.shl
                    local.tee 3
                    i32.const 1049568
                    i32.add
                    i32.load
                    local.tee 2
                    i32.const 8
                    i32.add
                    i32.load
                    local.tee 1
                    local.get 3
                    i32.const 1049560
                    i32.add
                    local.tee 3
                    i32.ne
                    if  ;; label = @9
                      local.get 1
                      local.get 3
                      i32.store offset=12
                      local.get 3
                      local.get 1
                      i32.store offset=8
                      br 1 (;@8;)
                    end
                    i32.const 1049824
                    i32.const 1049824
                    i32.load
                    i32.const -2
                    local.get 0
                    i32.rotl
                    i32.and
                    i32.store
                  end
                  local.get 2
                  local.get 4
                  call 36
                  local.get 2
                  local.get 4
                  call 39
                  local.tee 5
                  local.get 0
                  i32.const 3
                  i32.shl
                  local.get 4
                  i32.sub
                  local.tee 4
                  call 37
                  i32.const 1049832
                  i32.load
                  local.tee 3
                  if  ;; label = @8
                    local.get 3
                    i32.const -8
                    i32.and
                    i32.const 1049560
                    i32.add
                    local.set 0
                    i32.const 1049840
                    i32.load
                    local.set 1
                    block (result i32)  ;; label = @9
                      i32.const 1049824
                      i32.load
                      local.tee 6
                      i32.const 1
                      local.get 3
                      i32.const 3
                      i32.shr_u
                      i32.shl
                      local.tee 3
                      i32.and
                      if  ;; label = @10
                        local.get 0
                        i32.load offset=8
                        br 1 (;@9;)
                      end
                      i32.const 1049824
                      local.get 3
                      local.get 6
                      i32.or
                      i32.store
                      local.get 0
                    end
                    local.set 3
                    local.get 0
                    local.get 1
                    i32.store offset=8
                    local.get 3
                    local.get 1
                    i32.store offset=12
                    local.get 1
                    local.get 0
                    i32.store offset=12
                    local.get 1
                    local.get 3
                    i32.store offset=8
                  end
                  i32.const 1049840
                  local.get 5
                  i32.store
                  i32.const 1049832
                  local.get 4
                  i32.store
                  local.get 2
                  call 41
                  local.set 1
                  br 6 (;@1;)
                end
                i32.const 1049824
                local.get 6
                local.get 7
                i32.or
                i32.store
                local.get 0
              end
              local.set 6
              local.get 0
              local.get 3
              i32.store offset=8
              local.get 6
              local.get 3
              i32.store offset=12
              local.get 3
              local.get 0
              i32.store offset=12
              local.get 3
              local.get 6
              i32.store offset=8
            end
            i32.const 1049840
            local.get 5
            i32.store
            i32.const 1049832
            local.get 1
            i32.store
            br 1 (;@3;)
          end
          local.get 2
          local.get 1
          local.get 4
          i32.add
          call 35
        end
        local.get 2
        call 41
        local.tee 1
        br_if 1 (;@1;)
      end
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 4
                      i32.const 1049832
                      i32.load
                      local.tee 1
                      i32.gt_u
                      if  ;; label = @10
                        i32.const 1049836
                        i32.load
                        local.tee 0
                        local.get 4
                        i32.gt_u
                        br_if 2 (;@8;)
                        i32.const 8
                        i32.const 8
                        call 26
                        local.get 4
                        i32.add
                        i32.const 20
                        i32.const 8
                        call 26
                        i32.add
                        i32.const 16
                        i32.const 8
                        call 26
                        i32.add
                        i32.const 65536
                        call 26
                        local.tee 1
                        i32.const 16
                        i32.shr_u
                        memory.grow
                        local.set 0
                        local.get 8
                        i32.const 0
                        i32.store offset=8
                        local.get 8
                        i32.const 0
                        local.get 1
                        i32.const -65536
                        i32.and
                        local.get 0
                        i32.const -1
                        i32.eq
                        local.tee 1
                        select
                        i32.store offset=4
                        local.get 8
                        i32.const 0
                        local.get 0
                        i32.const 16
                        i32.shl
                        local.get 1
                        select
                        i32.store
                        local.get 8
                        i32.load
                        local.tee 1
                        br_if 1 (;@9;)
                        i32.const 0
                        local.set 1
                        br 9 (;@1;)
                      end
                      i32.const 1049840
                      i32.load
                      local.set 0
                      i32.const 16
                      i32.const 8
                      call 26
                      local.get 1
                      local.get 4
                      i32.sub
                      local.tee 1
                      i32.gt_u
                      if  ;; label = @10
                        i32.const 1049840
                        i32.const 0
                        i32.store
                        i32.const 1049832
                        i32.load
                        local.set 1
                        i32.const 1049832
                        i32.const 0
                        i32.store
                        local.get 0
                        local.get 1
                        call 35
                        local.get 0
                        call 41
                        local.set 1
                        br 9 (;@1;)
                      end
                      local.get 0
                      local.get 4
                      call 39
                      local.set 2
                      i32.const 1049832
                      local.get 1
                      i32.store
                      i32.const 1049840
                      local.get 2
                      i32.store
                      local.get 2
                      local.get 1
                      call 37
                      local.get 0
                      local.get 4
                      call 36
                      local.get 0
                      call 41
                      local.set 1
                      br 8 (;@1;)
                    end
                    local.get 8
                    i32.load offset=8
                    local.set 5
                    i32.const 1049848
                    local.get 8
                    i32.load offset=4
                    local.tee 3
                    i32.const 1049848
                    i32.load
                    i32.add
                    local.tee 0
                    i32.store
                    i32.const 1049852
                    i32.const 1049852
                    i32.load
                    local.tee 2
                    local.get 0
                    local.get 0
                    local.get 2
                    i32.lt_u
                    select
                    i32.store
                    block  ;; label = @9
                      block  ;; label = @10
                        i32.const 1049844
                        i32.load
                        if  ;; label = @11
                          i32.const 1049544
                          local.set 0
                          loop  ;; label = @12
                            local.get 0
                            call 46
                            local.get 1
                            i32.eq
                            br_if 2 (;@10;)
                            local.get 0
                            i32.load offset=8
                            local.tee 0
                            br_if 0 (;@12;)
                          end
                          br 2 (;@9;)
                        end
                        i32.const 1049860
                        i32.load
                        local.tee 0
                        i32.eqz
                        local.get 0
                        local.get 1
                        i32.gt_u
                        i32.or
                        br_if 3 (;@7;)
                        br 7 (;@3;)
                      end
                      local.get 0
                      call 44
                      br_if 0 (;@9;)
                      local.get 0
                      call 45
                      local.get 5
                      i32.ne
                      br_if 0 (;@9;)
                      local.get 0
                      i32.load
                      local.tee 2
                      i32.const 1049844
                      i32.load
                      local.tee 6
                      i32.le_u
                      if (result i32)  ;; label = @10
                        local.get 2
                        local.get 0
                        i32.load offset=4
                        i32.add
                        local.get 6
                        i32.gt_u
                      else
                        i32.const 0
                      end
                      br_if 3 (;@6;)
                    end
                    i32.const 1049860
                    i32.const 1049860
                    i32.load
                    local.tee 0
                    local.get 1
                    local.get 0
                    local.get 1
                    i32.lt_u
                    select
                    i32.store
                    local.get 1
                    local.get 3
                    i32.add
                    local.set 2
                    i32.const 1049544
                    local.set 0
                    block  ;; label = @9
                      block  ;; label = @10
                        loop  ;; label = @11
                          local.get 2
                          local.get 0
                          i32.load
                          i32.ne
                          if  ;; label = @12
                            local.get 0
                            i32.load offset=8
                            local.tee 0
                            br_if 1 (;@11;)
                            br 2 (;@10;)
                          end
                        end
                        local.get 0
                        call 44
                        br_if 0 (;@10;)
                        local.get 0
                        call 45
                        local.get 5
                        i32.eq
                        br_if 1 (;@9;)
                      end
                      i32.const 1049844
                      i32.load
                      local.set 2
                      i32.const 1049544
                      local.set 0
                      block  ;; label = @10
                        loop  ;; label = @11
                          local.get 2
                          local.get 0
                          i32.load
                          i32.ge_u
                          if  ;; label = @12
                            local.get 0
                            call 46
                            local.get 2
                            i32.gt_u
                            br_if 2 (;@10;)
                          end
                          local.get 0
                          i32.load offset=8
                          local.tee 0
                          br_if 0 (;@11;)
                        end
                        i32.const 0
                        local.set 0
                      end
                      local.get 2
                      local.get 0
                      call 46
                      local.tee 15
                      i32.const 20
                      i32.const 8
                      call 26
                      local.tee 14
                      i32.sub
                      i32.const 23
                      i32.sub
                      local.tee 0
                      call 41
                      local.tee 6
                      i32.const 8
                      call 26
                      local.get 6
                      i32.sub
                      local.get 0
                      i32.add
                      local.tee 0
                      local.get 0
                      i32.const 16
                      i32.const 8
                      call 26
                      local.get 2
                      i32.add
                      i32.lt_u
                      select
                      local.tee 6
                      call 41
                      local.set 7
                      local.get 6
                      local.get 14
                      call 39
                      local.set 0
                      i32.const 8
                      i32.const 8
                      call 26
                      local.set 9
                      i32.const 20
                      i32.const 8
                      call 26
                      local.set 11
                      i32.const 16
                      i32.const 8
                      call 26
                      local.set 12
                      i32.const 1049844
                      local.get 1
                      local.get 1
                      call 41
                      local.tee 10
                      i32.const 8
                      call 26
                      local.get 10
                      i32.sub
                      local.tee 13
                      call 39
                      local.tee 10
                      i32.store
                      i32.const 1049836
                      local.get 3
                      i32.const 8
                      i32.add
                      local.get 12
                      local.get 9
                      local.get 11
                      i32.add
                      i32.add
                      local.get 13
                      i32.add
                      i32.sub
                      local.tee 9
                      i32.store
                      local.get 10
                      local.get 9
                      i32.const 1
                      i32.or
                      i32.store offset=4
                      i32.const 8
                      i32.const 8
                      call 26
                      local.set 11
                      i32.const 20
                      i32.const 8
                      call 26
                      local.set 12
                      i32.const 16
                      i32.const 8
                      call 26
                      local.set 13
                      local.get 10
                      local.get 9
                      call 39
                      local.get 13
                      local.get 12
                      local.get 11
                      i32.const 8
                      i32.sub
                      i32.add
                      i32.add
                      i32.store offset=4
                      i32.const 1049856
                      i32.const 2097152
                      i32.store
                      local.get 6
                      local.get 14
                      call 36
                      i32.const 1049544
                      i64.load align=4
                      local.set 16
                      local.get 7
                      i32.const 8
                      i32.add
                      i32.const 1049552
                      i64.load align=4
                      i64.store align=4
                      local.get 7
                      local.get 16
                      i64.store align=4
                      i32.const 1049556
                      local.get 5
                      i32.store
                      i32.const 1049548
                      local.get 3
                      i32.store
                      i32.const 1049544
                      local.get 1
                      i32.store
                      i32.const 1049552
                      local.get 7
                      i32.store
                      loop  ;; label = @10
                        local.get 0
                        i32.const 4
                        call 39
                        local.get 0
                        i32.const 7
                        i32.store offset=4
                        local.tee 0
                        i32.const 4
                        i32.add
                        local.get 15
                        i32.lt_u
                        br_if 0 (;@10;)
                      end
                      local.get 2
                      local.get 6
                      i32.eq
                      br_if 7 (;@2;)
                      local.get 2
                      local.get 6
                      local.get 2
                      i32.sub
                      local.tee 0
                      local.get 2
                      local.get 0
                      call 39
                      call 38
                      local.get 0
                      i32.const 256
                      i32.ge_u
                      if  ;; label = @10
                        local.get 2
                        local.get 0
                        call 17
                        br 8 (;@2;)
                      end
                      local.get 0
                      i32.const -8
                      i32.and
                      i32.const 1049560
                      i32.add
                      local.set 1
                      block (result i32)  ;; label = @10
                        i32.const 1049824
                        i32.load
                        local.tee 3
                        i32.const 1
                        local.get 0
                        i32.const 3
                        i32.shr_u
                        i32.shl
                        local.tee 0
                        i32.and
                        if  ;; label = @11
                          local.get 1
                          i32.load offset=8
                          br 1 (;@10;)
                        end
                        i32.const 1049824
                        local.get 0
                        local.get 3
                        i32.or
                        i32.store
                        local.get 1
                      end
                      local.set 0
                      local.get 1
                      local.get 2
                      i32.store offset=8
                      local.get 0
                      local.get 2
                      i32.store offset=12
                      local.get 2
                      local.get 1
                      i32.store offset=12
                      local.get 2
                      local.get 0
                      i32.store offset=8
                      br 7 (;@2;)
                    end
                    local.get 0
                    i32.load
                    local.set 5
                    local.get 0
                    local.get 1
                    i32.store
                    local.get 0
                    local.get 0
                    i32.load offset=4
                    local.get 3
                    i32.add
                    i32.store offset=4
                    local.get 1
                    call 41
                    local.tee 0
                    i32.const 8
                    call 26
                    local.set 2
                    local.get 5
                    call 41
                    local.tee 3
                    i32.const 8
                    call 26
                    local.set 6
                    local.get 1
                    local.get 2
                    local.get 0
                    i32.sub
                    i32.add
                    local.tee 2
                    local.get 4
                    call 39
                    local.set 1
                    local.get 2
                    local.get 4
                    call 36
                    local.get 5
                    local.get 6
                    local.get 3
                    i32.sub
                    i32.add
                    local.tee 0
                    local.get 2
                    local.get 4
                    i32.add
                    i32.sub
                    local.set 4
                    i32.const 1049844
                    i32.load
                    local.get 0
                    i32.ne
                    if  ;; label = @9
                      local.get 0
                      i32.const 1049840
                      i32.load
                      i32.eq
                      br_if 4 (;@5;)
                      local.get 0
                      i32.load offset=4
                      i32.const 3
                      i32.and
                      i32.const 1
                      i32.ne
                      br_if 5 (;@4;)
                      block  ;; label = @10
                        local.get 0
                        call 30
                        local.tee 3
                        i32.const 256
                        i32.ge_u
                        if  ;; label = @11
                          local.get 0
                          call 16
                          br 1 (;@10;)
                        end
                        local.get 0
                        i32.const 12
                        i32.add
                        i32.load
                        local.tee 5
                        local.get 0
                        i32.const 8
                        i32.add
                        i32.load
                        local.tee 6
                        i32.ne
                        if  ;; label = @11
                          local.get 6
                          local.get 5
                          i32.store offset=12
                          local.get 5
                          local.get 6
                          i32.store offset=8
                          br 1 (;@10;)
                        end
                        i32.const 1049824
                        i32.const 1049824
                        i32.load
                        i32.const -2
                        local.get 3
                        i32.const 3
                        i32.shr_u
                        i32.rotl
                        i32.and
                        i32.store
                      end
                      local.get 3
                      local.get 4
                      i32.add
                      local.set 4
                      local.get 0
                      local.get 3
                      call 39
                      local.set 0
                      br 5 (;@4;)
                    end
                    i32.const 1049844
                    local.get 1
                    i32.store
                    i32.const 1049836
                    i32.const 1049836
                    i32.load
                    local.get 4
                    i32.add
                    local.tee 0
                    i32.store
                    local.get 1
                    local.get 0
                    i32.const 1
                    i32.or
                    i32.store offset=4
                    local.get 2
                    call 41
                    local.set 1
                    br 7 (;@1;)
                  end
                  i32.const 1049836
                  local.get 0
                  local.get 4
                  i32.sub
                  local.tee 1
                  i32.store
                  i32.const 1049844
                  i32.const 1049844
                  i32.load
                  local.tee 0
                  local.get 4
                  call 39
                  local.tee 2
                  i32.store
                  local.get 2
                  local.get 1
                  i32.const 1
                  i32.or
                  i32.store offset=4
                  local.get 0
                  local.get 4
                  call 36
                  local.get 0
                  call 41
                  local.set 1
                  br 6 (;@1;)
                end
                i32.const 1049860
                local.get 1
                i32.store
                br 3 (;@3;)
              end
              local.get 0
              local.get 0
              i32.load offset=4
              local.get 3
              i32.add
              i32.store offset=4
              i32.const 1049836
              i32.load
              local.get 3
              i32.add
              local.set 1
              i32.const 1049844
              i32.load
              local.tee 0
              local.get 0
              call 41
              local.tee 0
              i32.const 8
              call 26
              local.get 0
              i32.sub
              local.tee 2
              call 39
              local.set 0
              i32.const 1049836
              local.get 1
              local.get 2
              i32.sub
              local.tee 1
              i32.store
              i32.const 1049844
              local.get 0
              i32.store
              local.get 0
              local.get 1
              i32.const 1
              i32.or
              i32.store offset=4
              i32.const 8
              i32.const 8
              call 26
              local.set 2
              i32.const 20
              i32.const 8
              call 26
              local.set 3
              i32.const 16
              i32.const 8
              call 26
              local.set 5
              local.get 0
              local.get 1
              call 39
              local.get 5
              local.get 3
              local.get 2
              i32.const 8
              i32.sub
              i32.add
              i32.add
              i32.store offset=4
              i32.const 1049856
              i32.const 2097152
              i32.store
              br 3 (;@2;)
            end
            i32.const 1049840
            local.get 1
            i32.store
            i32.const 1049832
            i32.const 1049832
            i32.load
            local.get 4
            i32.add
            local.tee 0
            i32.store
            local.get 1
            local.get 0
            call 37
            local.get 2
            call 41
            local.set 1
            br 3 (;@1;)
          end
          local.get 1
          local.get 4
          local.get 0
          call 38
          local.get 4
          i32.const 256
          i32.ge_u
          if  ;; label = @4
            local.get 1
            local.get 4
            call 17
            local.get 2
            call 41
            local.set 1
            br 3 (;@1;)
          end
          local.get 4
          i32.const -8
          i32.and
          i32.const 1049560
          i32.add
          local.set 0
          block (result i32)  ;; label = @4
            i32.const 1049824
            i32.load
            local.tee 3
            i32.const 1
            local.get 4
            i32.const 3
            i32.shr_u
            i32.shl
            local.tee 5
            i32.and
            if  ;; label = @5
              local.get 0
              i32.load offset=8
              br 1 (;@4;)
            end
            i32.const 1049824
            local.get 3
            local.get 5
            i32.or
            i32.store
            local.get 0
          end
          local.set 3
          local.get 0
          local.get 1
          i32.store offset=8
          local.get 3
          local.get 1
          i32.store offset=12
          local.get 1
          local.get 0
          i32.store offset=12
          local.get 1
          local.get 3
          i32.store offset=8
          local.get 2
          call 41
          local.set 1
          br 2 (;@1;)
        end
        i32.const 1049864
        i32.const 4095
        i32.store
        i32.const 1049556
        local.get 5
        i32.store
        i32.const 1049548
        local.get 3
        i32.store
        i32.const 1049544
        local.get 1
        i32.store
        i32.const 1049572
        i32.const 1049560
        i32.store
        i32.const 1049580
        i32.const 1049568
        i32.store
        i32.const 1049568
        i32.const 1049560
        i32.store
        i32.const 1049588
        i32.const 1049576
        i32.store
        i32.const 1049576
        i32.const 1049568
        i32.store
        i32.const 1049596
        i32.const 1049584
        i32.store
        i32.const 1049584
        i32.const 1049576
        i32.store
        i32.const 1049604
        i32.const 1049592
        i32.store
        i32.const 1049592
        i32.const 1049584
        i32.store
        i32.const 1049612
        i32.const 1049600
        i32.store
        i32.const 1049600
        i32.const 1049592
        i32.store
        i32.const 1049620
        i32.const 1049608
        i32.store
        i32.const 1049608
        i32.const 1049600
        i32.store
        i32.const 1049628
        i32.const 1049616
        i32.store
        i32.const 1049616
        i32.const 1049608
        i32.store
        i32.const 1049636
        i32.const 1049624
        i32.store
        i32.const 1049624
        i32.const 1049616
        i32.store
        i32.const 1049632
        i32.const 1049624
        i32.store
        i32.const 1049644
        i32.const 1049632
        i32.store
        i32.const 1049640
        i32.const 1049632
        i32.store
        i32.const 1049652
        i32.const 1049640
        i32.store
        i32.const 1049648
        i32.const 1049640
        i32.store
        i32.const 1049660
        i32.const 1049648
        i32.store
        i32.const 1049656
        i32.const 1049648
        i32.store
        i32.const 1049668
        i32.const 1049656
        i32.store
        i32.const 1049664
        i32.const 1049656
        i32.store
        i32.const 1049676
        i32.const 1049664
        i32.store
        i32.const 1049672
        i32.const 1049664
        i32.store
        i32.const 1049684
        i32.const 1049672
        i32.store
        i32.const 1049680
        i32.const 1049672
        i32.store
        i32.const 1049692
        i32.const 1049680
        i32.store
        i32.const 1049688
        i32.const 1049680
        i32.store
        i32.const 1049700
        i32.const 1049688
        i32.store
        i32.const 1049708
        i32.const 1049696
        i32.store
        i32.const 1049696
        i32.const 1049688
        i32.store
        i32.const 1049716
        i32.const 1049704
        i32.store
        i32.const 1049704
        i32.const 1049696
        i32.store
        i32.const 1049724
        i32.const 1049712
        i32.store
        i32.const 1049712
        i32.const 1049704
        i32.store
        i32.const 1049732
        i32.const 1049720
        i32.store
        i32.const 1049720
        i32.const 1049712
        i32.store
        i32.const 1049740
        i32.const 1049728
        i32.store
        i32.const 1049728
        i32.const 1049720
        i32.store
        i32.const 1049748
        i32.const 1049736
        i32.store
        i32.const 1049736
        i32.const 1049728
        i32.store
        i32.const 1049756
        i32.const 1049744
        i32.store
        i32.const 1049744
        i32.const 1049736
        i32.store
        i32.const 1049764
        i32.const 1049752
        i32.store
        i32.const 1049752
        i32.const 1049744
        i32.store
        i32.const 1049772
        i32.const 1049760
        i32.store
        i32.const 1049760
        i32.const 1049752
        i32.store
        i32.const 1049780
        i32.const 1049768
        i32.store
        i32.const 1049768
        i32.const 1049760
        i32.store
        i32.const 1049788
        i32.const 1049776
        i32.store
        i32.const 1049776
        i32.const 1049768
        i32.store
        i32.const 1049796
        i32.const 1049784
        i32.store
        i32.const 1049784
        i32.const 1049776
        i32.store
        i32.const 1049804
        i32.const 1049792
        i32.store
        i32.const 1049792
        i32.const 1049784
        i32.store
        i32.const 1049812
        i32.const 1049800
        i32.store
        i32.const 1049800
        i32.const 1049792
        i32.store
        i32.const 1049820
        i32.const 1049808
        i32.store
        i32.const 1049808
        i32.const 1049800
        i32.store
        i32.const 1049816
        i32.const 1049808
        i32.store
        i32.const 8
        i32.const 8
        call 26
        local.set 2
        i32.const 20
        i32.const 8
        call 26
        local.set 5
        i32.const 16
        i32.const 8
        call 26
        local.set 6
        i32.const 1049844
        local.get 1
        local.get 1
        call 41
        local.tee 0
        i32.const 8
        call 26
        local.get 0
        i32.sub
        local.tee 1
        call 39
        local.tee 0
        i32.store
        i32.const 1049836
        local.get 3
        i32.const 8
        i32.add
        local.get 6
        local.get 2
        local.get 5
        i32.add
        i32.add
        local.get 1
        i32.add
        i32.sub
        local.tee 1
        i32.store
        local.get 0
        local.get 1
        i32.const 1
        i32.or
        i32.store offset=4
        i32.const 8
        i32.const 8
        call 26
        local.set 2
        i32.const 20
        i32.const 8
        call 26
        local.set 3
        i32.const 16
        i32.const 8
        call 26
        local.set 5
        local.get 0
        local.get 1
        call 39
        local.get 5
        local.get 3
        local.get 2
        i32.const 8
        i32.sub
        i32.add
        i32.add
        i32.store offset=4
        i32.const 1049856
        i32.const 2097152
        i32.store
      end
      i32.const 0
      local.set 1
      i32.const 1049836
      i32.load
      local.tee 0
      local.get 4
      i32.le_u
      br_if 0 (;@1;)
      i32.const 1049836
      local.get 0
      local.get 4
      i32.sub
      local.tee 1
      i32.store
      i32.const 1049844
      i32.const 1049844
      i32.load
      local.tee 0
      local.get 4
      call 39
      local.tee 2
      i32.store
      local.get 2
      local.get 1
      i32.const 1
      i32.or
      i32.store offset=4
      local.get 0
      local.get 4
      call 36
      local.get 0
      call 41
      local.set 1
    end
    local.get 8
    i32.const 16
    i32.add
    global.set 0
    local.get 1)
  (func (;15;) (type 0) (param i32 i32)
    (local i32 i32 i32 i32)
    local.get 0
    local.get 1
    call 39
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          call 32
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.set 3
          block  ;; label = @4
            local.get 0
            call 33
            i32.eqz
            if  ;; label = @5
              local.get 1
              local.get 3
              i32.add
              local.set 1
              local.get 0
              local.get 3
              call 40
              local.tee 0
              i32.const 1049840
              i32.load
              i32.ne
              br_if 1 (;@4;)
              local.get 2
              i32.load offset=4
              i32.const 3
              i32.and
              i32.const 3
              i32.ne
              br_if 2 (;@3;)
              i32.const 1049832
              local.get 1
              i32.store
              local.get 0
              local.get 1
              local.get 2
              call 38
              return
            end
            local.get 1
            local.get 3
            i32.add
            i32.const 16
            i32.add
            local.set 0
            br 2 (;@2;)
          end
          local.get 3
          i32.const 256
          i32.ge_u
          if  ;; label = @4
            local.get 0
            call 16
            br 1 (;@3;)
          end
          local.get 0
          i32.const 12
          i32.add
          i32.load
          local.tee 4
          local.get 0
          i32.const 8
          i32.add
          i32.load
          local.tee 5
          i32.ne
          if  ;; label = @4
            local.get 5
            local.get 4
            i32.store offset=12
            local.get 4
            local.get 5
            i32.store offset=8
            br 1 (;@3;)
          end
          i32.const 1049824
          i32.const 1049824
          i32.load
          i32.const -2
          local.get 3
          i32.const 3
          i32.shr_u
          i32.rotl
          i32.and
          i32.store
        end
        local.get 2
        call 31
        if  ;; label = @3
          local.get 0
          local.get 1
          local.get 2
          call 38
          br 2 (;@1;)
        end
        block  ;; label = @3
          i32.const 1049844
          i32.load
          local.get 2
          i32.ne
          if  ;; label = @4
            local.get 2
            i32.const 1049840
            i32.load
            i32.ne
            br_if 1 (;@3;)
            i32.const 1049840
            local.get 0
            i32.store
            i32.const 1049832
            i32.const 1049832
            i32.load
            local.get 1
            i32.add
            local.tee 1
            i32.store
            local.get 0
            local.get 1
            call 37
            return
          end
          i32.const 1049844
          local.get 0
          i32.store
          i32.const 1049836
          i32.const 1049836
          i32.load
          local.get 1
          i32.add
          local.tee 1
          i32.store
          local.get 0
          local.get 1
          i32.const 1
          i32.or
          i32.store offset=4
          local.get 0
          i32.const 1049840
          i32.load
          i32.ne
          br_if 1 (;@2;)
          i32.const 1049832
          i32.const 0
          i32.store
          i32.const 1049840
          i32.const 0
          i32.store
          return
        end
        local.get 2
        call 30
        local.tee 3
        local.get 1
        i32.add
        local.set 1
        block  ;; label = @3
          local.get 3
          i32.const 256
          i32.ge_u
          if  ;; label = @4
            local.get 2
            call 16
            br 1 (;@3;)
          end
          local.get 2
          i32.const 12
          i32.add
          i32.load
          local.tee 4
          local.get 2
          i32.const 8
          i32.add
          i32.load
          local.tee 2
          i32.ne
          if  ;; label = @4
            local.get 2
            local.get 4
            i32.store offset=12
            local.get 4
            local.get 2
            i32.store offset=8
            br 1 (;@3;)
          end
          i32.const 1049824
          i32.const 1049824
          i32.load
          i32.const -2
          local.get 3
          i32.const 3
          i32.shr_u
          i32.rotl
          i32.and
          i32.store
        end
        local.get 0
        local.get 1
        call 37
        local.get 0
        i32.const 1049840
        i32.load
        i32.ne
        br_if 1 (;@1;)
        i32.const 1049832
        local.get 1
        i32.store
      end
      return
    end
    local.get 1
    i32.const 256
    i32.ge_u
    if  ;; label = @1
      local.get 0
      local.get 1
      call 17
      return
    end
    local.get 1
    i32.const -8
    i32.and
    i32.const 1049560
    i32.add
    local.set 2
    block (result i32)  ;; label = @1
      i32.const 1049824
      i32.load
      local.tee 3
      i32.const 1
      local.get 1
      i32.const 3
      i32.shr_u
      i32.shl
      local.tee 1
      i32.and
      if  ;; label = @2
        local.get 2
        i32.load offset=8
        br 1 (;@1;)
      end
      i32.const 1049824
      local.get 1
      local.get 3
      i32.or
      i32.store
      local.get 2
    end
    local.set 1
    local.get 2
    local.get 0
    i32.store offset=8
    local.get 1
    local.get 0
    i32.store offset=12
    local.get 0
    local.get 2
    i32.store offset=12
    local.get 0
    local.get 1
    i32.store offset=8)
  (func (;16;) (type 4) (param i32)
    (local i32 i32 i32 i32 i32)
    local.get 0
    i32.load offset=24
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        local.get 0
        i32.load offset=12
        i32.eq
        if  ;; label = @3
          local.get 0
          i32.const 20
          i32.const 16
          local.get 0
          i32.const 20
          i32.add
          local.tee 1
          i32.load
          local.tee 3
          select
          i32.add
          i32.load
          local.tee 2
          br_if 1 (;@2;)
          i32.const 0
          local.set 1
          br 2 (;@1;)
        end
        local.get 0
        i32.load offset=8
        local.tee 2
        local.get 0
        i32.load offset=12
        local.tee 1
        i32.store offset=12
        local.get 1
        local.get 2
        i32.store offset=8
        br 1 (;@1;)
      end
      local.get 1
      local.get 0
      i32.const 16
      i32.add
      local.get 3
      select
      local.set 3
      loop  ;; label = @2
        local.get 3
        local.set 5
        local.get 2
        local.tee 1
        i32.const 20
        i32.add
        local.tee 3
        i32.load
        local.tee 2
        i32.eqz
        if  ;; label = @3
          local.get 1
          i32.const 16
          i32.add
          local.set 3
          local.get 1
          i32.load offset=16
          local.set 2
        end
        local.get 2
        br_if 0 (;@2;)
      end
      local.get 5
      i32.const 0
      i32.store
    end
    block  ;; label = @1
      local.get 4
      i32.eqz
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 0
        local.get 0
        i32.load offset=28
        i32.const 2
        i32.shl
        i32.const 1049416
        i32.add
        local.tee 2
        i32.load
        i32.ne
        if  ;; label = @3
          local.get 4
          i32.const 16
          i32.const 20
          local.get 4
          i32.load offset=16
          local.get 0
          i32.eq
          select
          i32.add
          local.get 1
          i32.store
          local.get 1
          br_if 1 (;@2;)
          br 2 (;@1;)
        end
        local.get 2
        local.get 1
        i32.store
        local.get 1
        br_if 0 (;@2;)
        i32.const 1049828
        i32.const 1049828
        i32.load
        i32.const -2
        local.get 0
        i32.load offset=28
        i32.rotl
        i32.and
        i32.store
        return
      end
      local.get 1
      local.get 4
      i32.store offset=24
      local.get 0
      i32.load offset=16
      local.tee 2
      if  ;; label = @2
        local.get 1
        local.get 2
        i32.store offset=16
        local.get 2
        local.get 1
        i32.store offset=24
      end
      local.get 0
      i32.const 20
      i32.add
      i32.load
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 20
      i32.add
      local.get 0
      i32.store
      local.get 0
      local.get 1
      i32.store offset=24
    end)
  (func (;17;) (type 0) (param i32 i32)
    (local i32 i32 i32 i32)
    local.get 0
    i64.const 0
    i64.store offset=16 align=4
    local.get 0
    block (result i32)  ;; label = @1
      i32.const 0
      local.get 1
      i32.const 256
      i32.lt_u
      br_if 0 (;@1;)
      drop
      i32.const 31
      local.get 1
      i32.const 16777215
      i32.gt_u
      br_if 0 (;@1;)
      drop
      local.get 1
      i32.const 6
      local.get 1
      i32.const 8
      i32.shr_u
      i32.clz
      local.tee 2
      i32.sub
      i32.shr_u
      i32.const 1
      i32.and
      local.get 2
      i32.const 1
      i32.shl
      i32.sub
      i32.const 62
      i32.add
    end
    local.tee 2
    i32.store offset=28
    local.get 2
    i32.const 2
    i32.shl
    i32.const 1049416
    i32.add
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            i32.const 1049828
            i32.load
            local.tee 4
            i32.const 1
            local.get 2
            i32.shl
            local.tee 5
            i32.and
            if  ;; label = @5
              local.get 3
              i32.load
              local.set 3
              local.get 2
              call 29
              local.set 2
              local.get 3
              call 30
              local.get 1
              i32.ne
              br_if 1 (;@4;)
              local.get 3
              local.set 2
              br 2 (;@3;)
            end
            i32.const 1049828
            local.get 4
            local.get 5
            i32.or
            i32.store
            local.get 3
            local.get 0
            i32.store
            br 3 (;@1;)
          end
          local.get 1
          local.get 2
          i32.shl
          local.set 4
          loop  ;; label = @4
            local.get 3
            local.get 4
            i32.const 29
            i32.shr_u
            i32.const 4
            i32.and
            i32.add
            i32.const 16
            i32.add
            local.tee 5
            i32.load
            local.tee 2
            i32.eqz
            br_if 2 (;@2;)
            local.get 4
            i32.const 1
            i32.shl
            local.set 4
            local.get 2
            local.tee 3
            call 30
            local.get 1
            i32.ne
            br_if 0 (;@4;)
          end
        end
        local.get 2
        i32.load offset=8
        local.tee 1
        local.get 0
        i32.store offset=12
        local.get 2
        local.get 0
        i32.store offset=8
        local.get 0
        local.get 2
        i32.store offset=12
        local.get 0
        local.get 1
        i32.store offset=8
        local.get 0
        i32.const 0
        i32.store offset=24
        return
      end
      local.get 5
      local.get 0
      i32.store
    end
    local.get 0
    local.get 3
    i32.store offset=24
    local.get 0
    local.get 0
    i32.store offset=8
    local.get 0
    local.get 0
    i32.store offset=12)
  (func (;18;) (type 7) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    i32.const 1049552
    i32.load
    local.tee 2
    if  ;; label = @1
      i32.const 1049544
      local.set 6
      loop  ;; label = @2
        local.get 2
        local.tee 1
        i32.load offset=8
        local.set 2
        local.get 1
        i32.load offset=4
        local.set 3
        local.get 1
        i32.load
        local.set 4
        local.get 1
        i32.const 12
        i32.add
        i32.load
        drop
        local.get 1
        local.set 6
        local.get 5
        i32.const 1
        i32.add
        local.set 5
        local.get 2
        br_if 0 (;@2;)
      end
    end
    i32.const 1049864
    local.get 5
    i32.const 4095
    local.get 5
    i32.const 4095
    i32.gt_u
    select
    i32.store
    local.get 8)
  (func (;19;) (type 4) (param i32)
    (local i32 i32 i32 i32 i32)
    local.get 0
    call 42
    local.tee 0
    local.get 0
    call 30
    local.tee 2
    call 39
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          call 32
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.set 3
          block  ;; label = @4
            local.get 0
            call 33
            i32.eqz
            if  ;; label = @5
              local.get 2
              local.get 3
              i32.add
              local.set 2
              local.get 0
              local.get 3
              call 40
              local.tee 0
              i32.const 1049840
              i32.load
              i32.ne
              br_if 1 (;@4;)
              local.get 1
              i32.load offset=4
              i32.const 3
              i32.and
              i32.const 3
              i32.ne
              br_if 2 (;@3;)
              i32.const 1049832
              local.get 2
              i32.store
              local.get 0
              local.get 2
              local.get 1
              call 38
              return
            end
            local.get 2
            local.get 3
            i32.add
            i32.const 16
            i32.add
            local.set 0
            br 2 (;@2;)
          end
          local.get 3
          i32.const 256
          i32.ge_u
          if  ;; label = @4
            local.get 0
            call 16
            br 1 (;@3;)
          end
          local.get 0
          i32.const 12
          i32.add
          i32.load
          local.tee 4
          local.get 0
          i32.const 8
          i32.add
          i32.load
          local.tee 5
          i32.ne
          if  ;; label = @4
            local.get 5
            local.get 4
            i32.store offset=12
            local.get 4
            local.get 5
            i32.store offset=8
            br 1 (;@3;)
          end
          i32.const 1049824
          i32.const 1049824
          i32.load
          i32.const -2
          local.get 3
          i32.const 3
          i32.shr_u
          i32.rotl
          i32.and
          i32.store
        end
        block  ;; label = @3
          local.get 1
          call 31
          if  ;; label = @4
            local.get 0
            local.get 2
            local.get 1
            call 38
            br 1 (;@3;)
          end
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                i32.const 1049844
                i32.load
                local.get 1
                i32.ne
                if  ;; label = @7
                  local.get 1
                  i32.const 1049840
                  i32.load
                  i32.ne
                  br_if 1 (;@6;)
                  i32.const 1049840
                  local.get 0
                  i32.store
                  i32.const 1049832
                  i32.const 1049832
                  i32.load
                  local.get 2
                  i32.add
                  local.tee 1
                  i32.store
                  local.get 0
                  local.get 1
                  call 37
                  return
                end
                i32.const 1049844
                local.get 0
                i32.store
                i32.const 1049836
                i32.const 1049836
                i32.load
                local.get 2
                i32.add
                local.tee 1
                i32.store
                local.get 0
                local.get 1
                i32.const 1
                i32.or
                i32.store offset=4
                local.get 0
                i32.const 1049840
                i32.load
                i32.eq
                br_if 1 (;@5;)
                br 2 (;@4;)
              end
              local.get 1
              call 30
              local.tee 3
              local.get 2
              i32.add
              local.set 2
              block  ;; label = @6
                local.get 3
                i32.const 256
                i32.ge_u
                if  ;; label = @7
                  local.get 1
                  call 16
                  br 1 (;@6;)
                end
                local.get 1
                i32.const 12
                i32.add
                i32.load
                local.tee 4
                local.get 1
                i32.const 8
                i32.add
                i32.load
                local.tee 1
                i32.ne
                if  ;; label = @7
                  local.get 1
                  local.get 4
                  i32.store offset=12
                  local.get 4
                  local.get 1
                  i32.store offset=8
                  br 1 (;@6;)
                end
                i32.const 1049824
                i32.const 1049824
                i32.load
                i32.const -2
                local.get 3
                i32.const 3
                i32.shr_u
                i32.rotl
                i32.and
                i32.store
              end
              local.get 0
              local.get 2
              call 37
              local.get 0
              i32.const 1049840
              i32.load
              i32.ne
              br_if 2 (;@3;)
              i32.const 1049832
              local.get 2
              i32.store
              br 3 (;@2;)
            end
            i32.const 1049832
            i32.const 0
            i32.store
            i32.const 1049840
            i32.const 0
            i32.store
          end
          i32.const 1049856
          i32.load
          local.get 1
          i32.ge_u
          br_if 1 (;@2;)
          i32.const 8
          i32.const 8
          call 26
          local.set 0
          i32.const 20
          i32.const 8
          call 26
          local.set 1
          i32.const 16
          i32.const 8
          call 26
          local.set 3
          i32.const 0
          i32.const 16
          i32.const 8
          call 26
          i32.const 2
          i32.shl
          i32.sub
          local.tee 2
          i32.const -65536
          local.get 3
          local.get 0
          local.get 1
          i32.add
          i32.add
          i32.sub
          i32.const -9
          i32.and
          i32.const 3
          i32.sub
          local.tee 0
          local.get 0
          local.get 2
          i32.gt_u
          select
          i32.eqz
          br_if 1 (;@2;)
          i32.const 1049844
          i32.load
          i32.eqz
          br_if 1 (;@2;)
          i32.const 8
          i32.const 8
          call 26
          local.set 0
          i32.const 20
          i32.const 8
          call 26
          local.set 1
          i32.const 16
          i32.const 8
          call 26
          local.set 2
          i32.const 0
          block  ;; label = @4
            i32.const 1049836
            i32.load
            local.tee 4
            local.get 2
            local.get 1
            local.get 0
            i32.const 8
            i32.sub
            i32.add
            i32.add
            local.tee 2
            i32.le_u
            br_if 0 (;@4;)
            i32.const 1049844
            i32.load
            local.set 1
            i32.const 1049544
            local.set 0
            block  ;; label = @5
              loop  ;; label = @6
                local.get 1
                local.get 0
                i32.load
                i32.ge_u
                if  ;; label = @7
                  local.get 0
                  call 46
                  local.get 1
                  i32.gt_u
                  br_if 2 (;@5;)
                end
                local.get 0
                i32.load offset=8
                local.tee 0
                br_if 0 (;@6;)
              end
              i32.const 0
              local.set 0
            end
            local.get 0
            call 44
            br_if 0 (;@4;)
            local.get 0
            i32.const 12
            i32.add
            i32.load
            drop
            br 0 (;@4;)
          end
          i32.const 0
          call 18
          i32.sub
          i32.ne
          br_if 1 (;@2;)
          i32.const 1049836
          i32.load
          i32.const 1049856
          i32.load
          i32.le_u
          br_if 1 (;@2;)
          i32.const 1049856
          i32.const -1
          i32.store
          return
        end
        local.get 2
        i32.const 256
        i32.lt_u
        br_if 1 (;@1;)
        local.get 0
        local.get 2
        call 17
        i32.const 1049864
        i32.const 1049864
        i32.load
        i32.const 1
        i32.sub
        local.tee 0
        i32.store
        local.get 0
        br_if 0 (;@2;)
        call 18
        drop
        return
      end
      return
    end
    local.get 2
    i32.const -8
    i32.and
    i32.const 1049560
    i32.add
    local.set 1
    block (result i32)  ;; label = @1
      i32.const 1049824
      i32.load
      local.tee 3
      i32.const 1
      local.get 2
      i32.const 3
      i32.shr_u
      i32.shl
      local.tee 2
      i32.and
      if  ;; label = @2
        local.get 1
        i32.load offset=8
        br 1 (;@1;)
      end
      i32.const 1049824
      local.get 2
      local.get 3
      i32.or
      i32.store
      local.get 1
    end
    local.set 3
    local.get 1
    local.get 0
    i32.store offset=8
    local.get 3
    local.get 0
    i32.store offset=12
    local.get 0
    local.get 1
    i32.store offset=12
    local.get 0
    local.get 3
    i32.store offset=8)
  (func (;20;) (type 0) (param i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 1
    global.set 0
    i32.const 1049388
    i32.load8_u
    if  ;; label = @1
      local.get 1
      i32.const 20
      i32.add
      i32.const 2
      i32.store
      local.get 1
      i32.const 28
      i32.add
      i32.const 1
      i32.store
      local.get 1
      i32.const 1048680
      i32.store offset=16
      local.get 1
      i32.const 0
      i32.store offset=8
      local.get 1
      i32.const 1
      i32.store offset=36
      local.get 1
      local.get 0
      i32.store offset=44
      local.get 1
      local.get 1
      i32.const 32
      i32.add
      i32.store offset=24
      local.get 1
      local.get 1
      i32.const 44
      i32.add
      i32.store offset=32
      local.get 1
      i32.const 8
      i32.add
      i32.const 1048720
      call 50
      unreachable
    end
    local.get 1
    i32.const 48
    i32.add
    global.set 0)
  (func (;21;) (type 0) (param i32 i32)
    (local i32 i32 i32 i64)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 1
    i32.load offset=4
    i32.eqz
    if  ;; label = @1
      local.get 1
      i32.load offset=12
      local.set 3
      local.get 2
      i32.const 16
      i32.add
      local.tee 4
      i32.const 0
      i32.store
      local.get 2
      i64.const 4294967296
      i64.store offset=8
      local.get 2
      local.get 2
      i32.const 8
      i32.add
      i32.store offset=20
      local.get 2
      i32.const 40
      i32.add
      local.get 3
      i32.const 16
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      i32.const 32
      i32.add
      local.get 3
      i32.const 8
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      local.get 3
      i64.load align=4
      i64.store offset=24
      local.get 2
      i32.const 20
      i32.add
      local.get 2
      i32.const 24
      i32.add
      call 53
      drop
      local.get 1
      i32.const 8
      i32.add
      local.get 4
      i32.load
      i32.store
      local.get 1
      local.get 2
      i64.load offset=8
      i64.store align=4
    end
    local.get 1
    i64.load align=4
    local.set 5
    local.get 1
    i64.const 4294967296
    i64.store align=4
    local.get 2
    i32.const 32
    i32.add
    local.tee 3
    local.get 1
    i32.const 8
    i32.add
    local.tee 1
    i32.load
    i32.store
    local.get 1
    i32.const 0
    i32.store
    local.get 2
    local.get 5
    i64.store offset=24
    i32.const 12
    i32.const 4
    call 1
    local.tee 1
    i32.eqz
    if  ;; label = @1
      i32.const 12
      i32.const 4
      call 47
      unreachable
    end
    local.get 1
    local.get 2
    i64.load offset=24
    i64.store align=4
    local.get 1
    i32.const 8
    i32.add
    local.get 3
    i32.load
    i32.store
    local.get 0
    i32.const 1048796
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func (;22;) (type 0) (param i32 i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 1
    i32.load offset=4
    i32.eqz
    if  ;; label = @1
      local.get 1
      i32.load offset=12
      local.set 3
      local.get 2
      i32.const 16
      i32.add
      local.tee 4
      i32.const 0
      i32.store
      local.get 2
      i64.const 4294967296
      i64.store offset=8
      local.get 2
      local.get 2
      i32.const 8
      i32.add
      i32.store offset=20
      local.get 2
      i32.const 40
      i32.add
      local.get 3
      i32.const 16
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      i32.const 32
      i32.add
      local.get 3
      i32.const 8
      i32.add
      i64.load align=4
      i64.store
      local.get 2
      local.get 3
      i64.load align=4
      i64.store offset=24
      local.get 2
      i32.const 20
      i32.add
      local.get 2
      i32.const 24
      i32.add
      call 53
      drop
      local.get 1
      i32.const 8
      i32.add
      local.get 4
      i32.load
      i32.store
      local.get 1
      local.get 2
      i64.load offset=8
      i64.store align=4
    end
    local.get 0
    i32.const 1048796
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func (;23;) (type 0) (param i32 i32)
    (local i32 i32)
    local.get 1
    i32.load offset=4
    local.set 2
    local.get 1
    i32.load
    local.set 3
    i32.const 8
    i32.const 4
    call 1
    local.tee 1
    i32.eqz
    if  ;; label = @1
      i32.const 8
      i32.const 4
      call 47
      unreachable
    end
    local.get 1
    local.get 2
    i32.store offset=4
    local.get 1
    local.get 3
    i32.store
    local.get 0
    i32.const 1048812
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;24;) (type 0) (param i32 i32)
    local.get 0
    i32.const 1048812
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;25;) (type 10) (param i32 i32 i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 5
    global.set 0
    i32.const 1049412
    i32.const 1049412
    i32.load
    local.tee 6
    i32.const 1
    i32.add
    i32.store
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 0
        i32.lt_s
        br_if 0 (;@2;)
        i32.const 1049868
        i32.const 1049868
        i32.load
        i32.const 1
        i32.add
        local.tee 6
        i32.store
        local.get 6
        i32.const 2
        i32.gt_u
        br_if 0 (;@2;)
        local.get 5
        local.get 4
        i32.store8 offset=24
        local.get 5
        local.get 3
        i32.store offset=20
        local.get 5
        local.get 2
        i32.store offset=16
        local.get 5
        i32.const 1048868
        i32.store offset=12
        local.get 5
        i32.const 1048600
        i32.store offset=8
        i32.const 1049396
        i32.load
        local.tee 2
        i32.const 0
        i32.lt_s
        br_if 0 (;@2;)
        i32.const 1049396
        local.get 2
        i32.const 1
        i32.add
        local.tee 2
        i32.store
        i32.const 1049396
        i32.const 1049404
        i32.load
        if (result i32)  ;; label = @3
          local.get 5
          local.get 0
          local.get 1
          i32.load offset=16
          call_indirect (type 0)
          local.get 5
          local.get 5
          i64.load
          i64.store offset=8
          i32.const 1049404
          i32.load
          local.get 5
          i32.const 8
          i32.add
          i32.const 1049408
          i32.load
          i32.load offset=20
          call_indirect (type 0)
          i32.const 1049396
          i32.load
        else
          local.get 2
        end
        i32.const 1
        i32.sub
        i32.store
        local.get 6
        i32.const 1
        i32.gt_u
        br_if 0 (;@2;)
        local.get 4
        br_if 1 (;@1;)
      end
      unreachable
    end
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.store offset=12
    local.get 2
    local.get 0
    i32.store offset=8
    unreachable)
  (func (;26;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add
    i32.const 1
    i32.sub
    i32.const 0
    local.get 1
    i32.sub
    i32.and)
  (func (;27;) (type 2) (param i32) (result i32)
    local.get 0
    i32.const 1
    i32.shl
    local.tee 0
    i32.const 0
    local.get 0
    i32.sub
    i32.or)
  (func (;28;) (type 2) (param i32) (result i32)
    i32.const 0
    local.get 0
    i32.sub
    local.get 0
    i32.and)
  (func (;29;) (type 2) (param i32) (result i32)
    i32.const 0
    i32.const 25
    local.get 0
    i32.const 1
    i32.shr_u
    i32.sub
    local.get 0
    i32.const 31
    i32.eq
    select)
  (func (;30;) (type 2) (param i32) (result i32)
    local.get 0
    i32.load offset=4
    i32.const -8
    i32.and)
  (func (;31;) (type 2) (param i32) (result i32)
    local.get 0
    i32.load8_u offset=4
    i32.const 2
    i32.and
    i32.const 1
    i32.shr_u)
  (func (;32;) (type 2) (param i32) (result i32)
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.and)
  (func (;33;) (type 2) (param i32) (result i32)
    local.get 0
    i32.load8_u offset=4
    i32.const 3
    i32.and
    i32.eqz)
  (func (;34;) (type 0) (param i32 i32)
    local.get 0
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.and
    local.get 1
    i32.or
    i32.const 2
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.tee 0
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.or
    i32.store offset=4)
  (func (;35;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 3
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.tee 0
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.or
    i32.store offset=4)
  (func (;36;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 3
    i32.or
    i32.store offset=4)
  (func (;37;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 1
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.get 1
    i32.store)
  (func (;38;) (type 6) (param i32 i32 i32)
    local.get 2
    local.get 2
    i32.load offset=4
    i32.const -2
    i32.and
    i32.store offset=4
    local.get 0
    local.get 1
    i32.const 1
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.get 1
    i32.store)
  (func (;39;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add)
  (func (;40;) (type 1) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.sub)
  (func (;41;) (type 2) (param i32) (result i32)
    local.get 0
    i32.const 8
    i32.add)
  (func (;42;) (type 2) (param i32) (result i32)
    local.get 0
    i32.const 8
    i32.sub)
  (func (;43;) (type 2) (param i32) (result i32)
    (local i32)
    local.get 0
    i32.load offset=16
    local.tee 1
    if (result i32)  ;; label = @1
      local.get 1
    else
      local.get 0
      i32.const 20
      i32.add
      i32.load
    end)
  (func (;44;) (type 2) (param i32) (result i32)
    local.get 0
    i32.load offset=12
    i32.const 1
    i32.and)
  (func (;45;) (type 2) (param i32) (result i32)
    local.get 0
    i32.load offset=12
    i32.const 1
    i32.shr_u)
  (func (;46;) (type 2) (param i32) (result i32)
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    i32.add)
  (func (;47;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 1049392
    i32.load
    local.tee 0
    i32.const 2
    local.get 0
    select
    call_indirect (type 0)
    unreachable)
  (func (;48;) (type 8)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    local.get 0
    i32.const 20
    i32.add
    i32.const 1
    i32.store
    local.get 0
    i32.const 28
    i32.add
    i32.const 0
    i32.store
    local.get 0
    i32.const 1048932
    i32.store offset=16
    local.get 0
    i32.const 1048884
    i32.store offset=24
    local.get 0
    i32.const 0
    i32.store offset=8
    local.get 0
    i32.const 8
    i32.add
    i32.const 1048940
    call 50
    unreachable)
  (func (;49;) (type 1) (param i32 i32) (result i32)
    local.get 0
    i32.load
    drop
    loop  ;; label = @1
      br 0 (;@1;)
    end
    unreachable)
  (func (;50;) (type 0) (param i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 1
    i32.store8 offset=24
    local.get 2
    local.get 1
    i32.store offset=20
    local.get 2
    local.get 0
    i32.store offset=16
    local.get 2
    i32.const 1048956
    i32.store offset=12
    local.get 2
    i32.const 1048956
    i32.store offset=8
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      local.get 2
      i32.const 8
      i32.add
      local.tee 0
      i32.load offset=12
      local.tee 2
      if  ;; label = @2
        local.get 0
        i32.load offset=8
        local.tee 3
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        local.get 2
        i32.store offset=8
        local.get 1
        local.get 0
        i32.store offset=4
        local.get 1
        local.get 3
        i32.store
        global.get 0
        i32.const 16
        i32.sub
        local.tee 0
        global.set 0
        local.get 0
        i32.const 8
        i32.add
        local.get 1
        i32.const 8
        i32.add
        i32.load
        i32.store
        local.get 0
        local.get 1
        i64.load align=4
        i64.store
        global.get 0
        i32.const 16
        i32.sub
        local.tee 1
        global.set 0
        local.get 0
        i32.load
        local.tee 2
        i32.const 20
        i32.add
        i32.load
        local.set 3
        block  ;; label = @3
          block (result i32)  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 2
                i32.const 12
                i32.add
                i32.load
                br_table 0 (;@6;) 1 (;@5;) 3 (;@3;)
              end
              local.get 3
              br_if 2 (;@3;)
              i32.const 0
              local.set 2
              i32.const 1048600
              br 1 (;@4;)
            end
            local.get 3
            br_if 1 (;@3;)
            local.get 2
            i32.load offset=8
            local.tee 3
            i32.load offset=4
            local.set 2
            local.get 3
            i32.load
          end
          local.set 3
          local.get 1
          local.get 2
          i32.store offset=4
          local.get 1
          local.get 3
          i32.store
          local.get 1
          i32.const 1048848
          local.get 0
          i32.load offset=4
          local.tee 1
          i32.load offset=8
          local.get 0
          i32.load offset=8
          local.get 1
          i32.load8_u offset=16
          call 25
          unreachable
        end
        local.get 1
        i32.const 0
        i32.store offset=4
        local.get 1
        local.get 2
        i32.store offset=12
        local.get 1
        i32.const 1048828
        local.get 0
        i32.load offset=4
        local.tee 1
        i32.load offset=8
        local.get 0
        i32.load offset=8
        local.get 1
        i32.load8_u offset=16
        call 25
        unreachable
      end
      i32.const 1048780
      call 51
      unreachable
    end
    i32.const 1048764
    call 51
    unreachable)
  (func (;51;) (type 4) (param i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i32.const 12
    i32.add
    i32.const 1
    i32.store
    local.get 1
    i32.const 20
    i32.add
    i32.const 0
    i32.store
    local.get 1
    i32.const 1048956
    i32.store offset=16
    local.get 1
    i32.const 0
    i32.store
    local.get 1
    i32.const 43
    i32.store offset=28
    local.get 1
    i32.const 1048600
    i32.store offset=24
    local.get 1
    local.get 1
    i32.const 24
    i32.add
    i32.store offset=8
    local.get 1
    local.get 0
    call 50
    unreachable)
  (func (;52;) (type 1) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64 i64)
    local.get 0
    i64.load32_u
    local.set 13
    global.get 0
    i32.const 48
    i32.sub
    local.tee 4
    global.set 0
    i32.const 39
    local.set 0
    block  ;; label = @1
      local.get 13
      i64.const 10000
      i64.lt_u
      if  ;; label = @2
        local.get 13
        local.set 14
        br 1 (;@1;)
      end
      loop  ;; label = @2
        local.get 4
        i32.const 9
        i32.add
        local.get 0
        i32.add
        local.tee 2
        i32.const 4
        i32.sub
        local.get 13
        local.get 13
        i64.const 10000
        i64.div_u
        local.tee 14
        i64.const 10000
        i64.mul
        i64.sub
        i32.wrap_i64
        local.tee 3
        i32.const 65535
        i32.and
        i32.const 100
        i32.div_u
        local.tee 5
        i32.const 1
        i32.shl
        i32.const 1048972
        i32.add
        i32.load16_u align=1
        i32.store16 align=1
        local.get 2
        i32.const 2
        i32.sub
        local.get 3
        local.get 5
        i32.const 100
        i32.mul
        i32.sub
        i32.const 65535
        i32.and
        i32.const 1
        i32.shl
        i32.const 1048972
        i32.add
        i32.load16_u align=1
        i32.store16 align=1
        local.get 0
        i32.const 4
        i32.sub
        local.set 0
        local.get 13
        i64.const 99999999
        i64.gt_u
        local.get 14
        local.set 13
        br_if 0 (;@2;)
      end
    end
    local.get 14
    i32.wrap_i64
    local.tee 2
    i32.const 99
    i32.gt_u
    if  ;; label = @1
      local.get 0
      i32.const 2
      i32.sub
      local.tee 0
      local.get 4
      i32.const 9
      i32.add
      i32.add
      local.get 14
      i32.wrap_i64
      local.tee 2
      local.get 2
      i32.const 65535
      i32.and
      i32.const 100
      i32.div_u
      local.tee 2
      i32.const 100
      i32.mul
      i32.sub
      i32.const 65535
      i32.and
      i32.const 1
      i32.shl
      i32.const 1048972
      i32.add
      i32.load16_u align=1
      i32.store16 align=1
    end
    block  ;; label = @1
      local.get 2
      i32.const 10
      i32.ge_u
      if  ;; label = @2
        local.get 0
        i32.const 2
        i32.sub
        local.tee 0
        local.get 4
        i32.const 9
        i32.add
        i32.add
        local.get 2
        i32.const 1
        i32.shl
        i32.const 1048972
        i32.add
        i32.load16_u align=1
        i32.store16 align=1
        br 1 (;@1;)
      end
      local.get 0
      i32.const 1
      i32.sub
      local.tee 0
      local.get 4
      i32.const 9
      i32.add
      i32.add
      local.get 2
      i32.const 48
      i32.add
      i32.store8
    end
    block (result i32)  ;; label = @1
      local.get 4
      i32.const 9
      i32.add
      local.get 0
      i32.add
      local.set 8
      i32.const 43
      i32.const 1114112
      local.get 1
      i32.load offset=24
      local.tee 3
      i32.const 1
      i32.and
      local.tee 2
      select
      local.set 5
      local.get 2
      i32.const 39
      local.get 0
      i32.sub
      local.tee 9
      i32.add
      local.set 2
      i32.const 1048956
      i32.const 0
      local.get 3
      i32.const 4
      i32.and
      select
      local.set 7
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.load offset=8
          i32.eqz
          if  ;; label = @4
            i32.const 1
            local.set 0
            local.get 1
            i32.load
            local.tee 3
            local.get 1
            i32.const 4
            i32.add
            i32.load
            local.tee 1
            local.get 5
            local.get 7
            call 54
            br_if 1 (;@3;)
            br 2 (;@2;)
          end
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  local.get 1
                  i32.const 12
                  i32.add
                  i32.load
                  local.tee 6
                  i32.lt_u
                  if  ;; label = @8
                    local.get 3
                    i32.const 8
                    i32.and
                    br_if 4 (;@4;)
                    local.get 6
                    local.get 2
                    i32.sub
                    local.tee 2
                    local.set 3
                    i32.const 1
                    local.get 1
                    i32.load8_u offset=32
                    local.tee 0
                    local.get 0
                    i32.const 3
                    i32.eq
                    select
                    i32.const 3
                    i32.and
                    local.tee 0
                    i32.const 1
                    i32.sub
                    br_table 1 (;@7;) 2 (;@6;) 3 (;@5;)
                  end
                  i32.const 1
                  local.set 0
                  local.get 1
                  i32.load
                  local.tee 3
                  local.get 1
                  i32.const 4
                  i32.add
                  i32.load
                  local.tee 1
                  local.get 5
                  local.get 7
                  call 54
                  br_if 4 (;@3;)
                  br 5 (;@2;)
                end
                i32.const 0
                local.set 3
                local.get 2
                local.set 0
                br 1 (;@5;)
              end
              local.get 2
              i32.const 1
              i32.shr_u
              local.set 0
              local.get 2
              i32.const 1
              i32.add
              i32.const 1
              i32.shr_u
              local.set 3
            end
            local.get 0
            i32.const 1
            i32.add
            local.set 0
            local.get 1
            i32.const 4
            i32.add
            i32.load
            local.set 2
            local.get 1
            i32.load offset=28
            local.set 6
            local.get 1
            i32.load
            local.set 1
            block  ;; label = @5
              loop  ;; label = @6
                local.get 0
                i32.const 1
                i32.sub
                local.tee 0
                i32.eqz
                br_if 1 (;@5;)
                local.get 1
                local.get 6
                local.get 2
                i32.load offset=16
                call_indirect (type 1)
                i32.eqz
                br_if 0 (;@6;)
              end
              i32.const 1
              br 4 (;@1;)
            end
            i32.const 1
            local.set 0
            local.get 6
            i32.const 1114112
            i32.eq
            br_if 1 (;@3;)
            local.get 1
            local.get 2
            local.get 5
            local.get 7
            call 54
            br_if 1 (;@3;)
            local.get 1
            local.get 8
            local.get 9
            local.get 2
            i32.load offset=12
            call_indirect (type 3)
            br_if 1 (;@3;)
            i32.const 0
            local.set 0
            block (result i32)  ;; label = @5
              loop  ;; label = @6
                local.get 3
                local.get 0
                local.get 3
                i32.eq
                br_if 1 (;@5;)
                drop
                local.get 0
                i32.const 1
                i32.add
                local.set 0
                local.get 1
                local.get 6
                local.get 2
                i32.load offset=16
                call_indirect (type 1)
                i32.eqz
                br_if 0 (;@6;)
              end
              local.get 0
              i32.const 1
              i32.sub
            end
            local.get 3
            i32.lt_u
            local.set 0
            br 1 (;@3;)
          end
          local.get 1
          i32.load offset=28
          local.set 11
          local.get 1
          i32.const 48
          i32.store offset=28
          local.get 1
          i32.load8_u offset=32
          local.set 12
          i32.const 1
          local.set 0
          local.get 1
          i32.const 1
          i32.store8 offset=32
          local.get 1
          i32.load
          local.tee 3
          local.get 1
          i32.const 4
          i32.add
          i32.load
          local.tee 10
          local.get 5
          local.get 7
          call 54
          br_if 0 (;@3;)
          local.get 6
          local.get 2
          i32.sub
          i32.const 1
          i32.add
          local.set 0
          block  ;; label = @4
            loop  ;; label = @5
              local.get 0
              i32.const 1
              i32.sub
              local.tee 0
              i32.eqz
              br_if 1 (;@4;)
              local.get 3
              i32.const 48
              local.get 10
              i32.load offset=16
              call_indirect (type 1)
              i32.eqz
              br_if 0 (;@5;)
            end
            i32.const 1
            br 3 (;@1;)
          end
          i32.const 1
          local.set 0
          local.get 3
          local.get 8
          local.get 9
          local.get 10
          i32.load offset=12
          call_indirect (type 3)
          br_if 0 (;@3;)
          local.get 1
          local.get 12
          i32.store8 offset=32
          local.get 1
          local.get 11
          i32.store offset=28
          i32.const 0
          br 2 (;@1;)
        end
        local.get 0
        br 1 (;@1;)
      end
      local.get 3
      local.get 8
      local.get 9
      local.get 1
      i32.load offset=12
      call_indirect (type 3)
    end
    local.get 4
    i32.const 48
    i32.add
    global.set 0)
  (func (;53;) (type 1) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 3
    i32.store8 offset=40
    local.get 2
    i64.const 137438953472
    i64.store offset=32
    local.get 2
    i32.const 0
    i32.store offset=24
    local.get 2
    i32.const 0
    i32.store offset=16
    local.get 2
    i32.const 1048576
    i32.store offset=12
    local.get 2
    local.get 0
    i32.store offset=8
    block (result i32)  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.load
          local.tee 10
          i32.eqz
          if  ;; label = @4
            local.get 1
            i32.const 20
            i32.add
            i32.load
            local.tee 0
            i32.eqz
            br_if 1 (;@3;)
            local.get 1
            i32.load offset=16
            local.set 3
            local.get 0
            i32.const 3
            i32.shl
            local.set 5
            local.get 0
            i32.const 1
            i32.sub
            i32.const 536870911
            i32.and
            i32.const 1
            i32.add
            local.set 7
            local.get 1
            i32.load offset=8
            local.set 0
            loop  ;; label = @5
              local.get 0
              i32.const 4
              i32.add
              i32.load
              local.tee 4
              if  ;; label = @6
                local.get 2
                i32.load offset=8
                local.get 0
                i32.load
                local.get 4
                local.get 2
                i32.load offset=12
                i32.load offset=12
                call_indirect (type 3)
                br_if 4 (;@2;)
              end
              local.get 3
              i32.load
              local.get 2
              i32.const 8
              i32.add
              local.get 3
              i32.const 4
              i32.add
              i32.load
              call_indirect (type 1)
              br_if 3 (;@2;)
              local.get 3
              i32.const 8
              i32.add
              local.set 3
              local.get 0
              i32.const 8
              i32.add
              local.set 0
              local.get 5
              i32.const 8
              i32.sub
              local.tee 5
              br_if 0 (;@5;)
            end
            br 1 (;@3;)
          end
          local.get 1
          i32.load offset=4
          local.tee 0
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          i32.const 5
          i32.shl
          local.set 11
          local.get 0
          i32.const 1
          i32.sub
          i32.const 134217727
          i32.and
          i32.const 1
          i32.add
          local.set 7
          local.get 1
          i32.load offset=8
          local.set 0
          loop  ;; label = @4
            local.get 0
            i32.const 4
            i32.add
            i32.load
            local.tee 3
            if  ;; label = @5
              local.get 2
              i32.load offset=8
              local.get 0
              i32.load
              local.get 3
              local.get 2
              i32.load offset=12
              i32.load offset=12
              call_indirect (type 3)
              br_if 3 (;@2;)
            end
            local.get 2
            local.get 5
            local.get 10
            i32.add
            local.tee 4
            i32.const 28
            i32.add
            i32.load8_u
            i32.store8 offset=40
            local.get 2
            local.get 4
            i32.const 20
            i32.add
            i64.load align=4
            i64.store offset=32
            local.get 4
            i32.const 16
            i32.add
            i32.load
            local.set 6
            local.get 1
            i32.load offset=16
            local.set 8
            i32.const 0
            local.set 9
            i32.const 0
            local.set 3
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 4
                  i32.const 12
                  i32.add
                  i32.load
                  i32.const 1
                  i32.sub
                  br_table 0 (;@7;) 2 (;@5;) 1 (;@6;)
                end
                local.get 6
                i32.const 3
                i32.shl
                local.get 8
                i32.add
                local.tee 12
                i32.const 4
                i32.add
                i32.load
                i32.const 16
                i32.ne
                br_if 1 (;@5;)
                local.get 12
                i32.load
                i32.load
                local.set 6
              end
              i32.const 1
              local.set 3
            end
            local.get 2
            local.get 6
            i32.store offset=20
            local.get 2
            local.get 3
            i32.store offset=16
            local.get 4
            i32.const 8
            i32.add
            i32.load
            local.set 3
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 4
                  i32.const 4
                  i32.add
                  i32.load
                  i32.const 1
                  i32.sub
                  br_table 0 (;@7;) 2 (;@5;) 1 (;@6;)
                end
                local.get 3
                i32.const 3
                i32.shl
                local.get 8
                i32.add
                local.tee 6
                i32.const 4
                i32.add
                i32.load
                i32.const 16
                i32.ne
                br_if 1 (;@5;)
                local.get 6
                i32.load
                i32.load
                local.set 3
              end
              i32.const 1
              local.set 9
            end
            local.get 2
            local.get 3
            i32.store offset=28
            local.get 2
            local.get 9
            i32.store offset=24
            local.get 8
            local.get 4
            i32.load
            i32.const 3
            i32.shl
            i32.add
            local.tee 3
            i32.load
            local.get 2
            i32.const 8
            i32.add
            local.get 3
            i32.load offset=4
            call_indirect (type 1)
            br_if 2 (;@2;)
            local.get 0
            i32.const 8
            i32.add
            local.set 0
            local.get 11
            local.get 5
            i32.const 32
            i32.add
            local.tee 5
            i32.ne
            br_if 0 (;@4;)
          end
        end
        local.get 1
        i32.const 12
        i32.add
        i32.load
        local.get 7
        i32.gt_u
        if  ;; label = @3
          local.get 2
          i32.load offset=8
          local.get 1
          i32.load offset=8
          local.get 7
          i32.const 3
          i32.shl
          i32.add
          local.tee 0
          i32.load
          local.get 0
          i32.load offset=4
          local.get 2
          i32.load offset=12
          i32.load offset=12
          call_indirect (type 3)
          br_if 1 (;@2;)
        end
        i32.const 0
        br 1 (;@1;)
      end
      i32.const 1
    end
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func (;54;) (type 11) (param i32 i32 i32 i32) (result i32)
    block  ;; label = @1
      block (result i32)  ;; label = @2
        local.get 2
        i32.const 1114112
        i32.ne
        if  ;; label = @3
          i32.const 1
          local.get 0
          local.get 2
          local.get 1
          i32.load offset=16
          call_indirect (type 1)
          br_if 1 (;@2;)
          drop
        end
        local.get 3
        br_if 1 (;@1;)
        i32.const 0
      end
      return
    end
    local.get 0
    local.get 3
    i32.const 0
    local.get 1
    i32.load offset=12
    call_indirect (type 3))
  (func (;55;) (type 3) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      local.get 2
      local.tee 4
      i32.const 15
      i32.le_u
      if  ;; label = @2
        local.get 0
        local.set 2
        br 1 (;@1;)
      end
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 3
      i32.add
      local.set 5
      local.get 3
      if  ;; label = @2
        local.get 0
        local.set 2
        local.get 1
        local.set 6
        loop  ;; label = @3
          local.get 2
          local.get 6
          i32.load8_u
          i32.store8
          local.get 6
          i32.const 1
          i32.add
          local.set 6
          local.get 2
          i32.const 1
          i32.add
          local.tee 2
          local.get 5
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 5
      local.get 4
      local.get 3
      i32.sub
      local.tee 8
      i32.const -4
      i32.and
      local.tee 7
      i32.add
      local.set 2
      block  ;; label = @2
        local.get 1
        local.get 3
        i32.add
        local.tee 3
        i32.const 3
        i32.and
        local.tee 4
        if  ;; label = @3
          local.get 7
          i32.const 0
          i32.le_s
          br_if 1 (;@2;)
          local.get 3
          i32.const -4
          i32.and
          local.tee 6
          i32.const 4
          i32.add
          local.set 1
          i32.const 0
          local.get 4
          i32.const 3
          i32.shl
          local.tee 9
          i32.sub
          i32.const 24
          i32.and
          local.set 4
          local.get 6
          i32.load
          local.set 6
          loop  ;; label = @4
            local.get 5
            local.get 6
            local.get 9
            i32.shr_u
            local.get 1
            i32.load
            local.tee 6
            local.get 4
            i32.shl
            i32.or
            i32.store
            local.get 1
            i32.const 4
            i32.add
            local.set 1
            local.get 5
            i32.const 4
            i32.add
            local.tee 5
            local.get 2
            i32.lt_u
            br_if 0 (;@4;)
          end
          br 1 (;@2;)
        end
        local.get 7
        i32.const 0
        i32.le_s
        br_if 0 (;@2;)
        local.get 3
        local.set 1
        loop  ;; label = @3
          local.get 5
          local.get 1
          i32.load
          i32.store
          local.get 1
          i32.const 4
          i32.add
          local.set 1
          local.get 5
          i32.const 4
          i32.add
          local.tee 5
          local.get 2
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 8
      i32.const 3
      i32.and
      local.set 4
      local.get 3
      local.get 7
      i32.add
      local.set 1
    end
    local.get 4
    if  ;; label = @1
      local.get 2
      local.get 4
      i32.add
      local.set 3
      loop  ;; label = @2
        local.get 2
        local.get 1
        i32.load8_u
        i32.store8
        local.get 1
        i32.const 1
        i32.add
        local.set 1
        local.get 2
        i32.const 1
        i32.add
        local.tee 2
        local.get 3
        i32.lt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func (;56;) (type 6) (param i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    local.get 2
    i32.store offset=8
    local.get 3
    local.get 1
    i32.store offset=4
    local.get 3
    local.get 0
    i32.store
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    local.get 0
    i32.const 8
    i32.add
    local.get 3
    i32.const 8
    i32.add
    i32.load
    i32.store
    local.get 0
    local.get 3
    i64.load align=4
    i64.store
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    local.get 0
    i64.load align=4
    i64.store offset=8
    local.get 1
    i32.const 8
    i32.add
    i32.const 1049172
    i32.const 0
    local.get 0
    i32.load offset=8
    i32.const 1
    call 25
    unreachable)
  (func (;57;) (type 0) (param i32 i32)
    local.get 1
    i32.load
    i32.eqz
    if  ;; label = @1
      unreachable
    end
    local.get 0
    i32.const 1049192
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;58;) (type 0) (param i32 i32)
    (local i32 i32)
    local.get 1
    i32.load
    local.set 2
    local.get 1
    i32.const 0
    i32.store
    block  ;; label = @1
      local.get 2
      if  ;; label = @2
        local.get 1
        i32.load offset=4
        local.set 3
        i32.const 8
        i32.const 4
        call 1
        local.tee 1
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        local.get 3
        i32.store offset=4
        local.get 1
        local.get 2
        i32.store
        local.get 0
        i32.const 1049192
        i32.store offset=4
        local.get 0
        local.get 1
        i32.store
        return
      end
      unreachable
    end
    i32.const 8
    i32.const 4
    call 47
    unreachable)
  (func (;59;) (type 2) (param i32) (result i32)
    (local i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 0
            i32.eqz
            if  ;; label = @5
              i32.const 1
              local.set 2
              br 1 (;@4;)
            end
            local.get 0
            i32.const 0
            i32.ge_s
            local.tee 1
            i32.eqz
            br_if 1 (;@3;)
            local.get 0
            local.get 1
            call 1
            local.tee 2
            i32.eqz
            br_if 2 (;@2;)
          end
          i32.const 12
          i32.const 4
          call 1
          local.tee 1
          i32.eqz
          br_if 2 (;@1;)
          local.get 1
          i32.const 0
          i32.store offset=8
          local.get 1
          local.get 0
          i32.store offset=4
          local.get 1
          local.get 2
          i32.store
          local.get 1
          return
        end
        call 48
        unreachable
      end
      local.get 0
      local.get 1
      call 47
      unreachable
    end
    i32.const 12
    i32.const 4
    call 47
    unreachable)
  (func (;60;) (type 4) (param i32)
    (local i32 i32)
    block  ;; label = @1
      local.get 0
      if  ;; label = @2
        local.get 0
        i32.load
        local.tee 1
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        i32.load offset=4
        local.get 0
        call 19
        if  ;; label = @3
          local.get 1
          call 19
        end
        return
      end
      i32.const 1049208
      i32.const 22
      i32.const 1049324
      call 56
      unreachable
    end
    i32.const 1049340
    i32.const 29
    i32.const 1049372
    call 56
    unreachable)
  (func (;61;) (type 8)
    nop)
  (table (;0;) 23 23 funcref)
  (memory (;0;) 17)
  (global (;0;) (mut i32) (i32.const 1048576))
  (global (;1;) i32 (i32.const 1049872))
  (global (;2;) i32 (i32.const 1049872))
  (export "memory" (memory 0))
  (export "interface_version_8" (func 0))
  (export "allocate" (func 59))
  (export "deallocate" (func 60))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (elem (;0;) (i32.const 1) func 52 20 6 11 9 10 7 2 3 8 21 22 23 24 4 49 6 4 6 58 57 3)
  (data (;0;) (i32.const 1048576) "\03\00\00\00\04\00\00\00\04\00\00\00\04\00\00\00\05\00\00\00\06\00\00\00called `Option::unwrap()` on a `None` valuememory allocation of  bytes failed\0a\00\00C\00\10\00\15\00\00\00X\00\10\00\0e\00\00\00library/std/src/alloc.rsx\00\10\00\18\00\00\00U\01\00\00\09\00\00\00library/std/src/panicking.rs\a0\00\10\00\1c\00\00\00>\02\00\00\0f\00\00\00\a0\00\10\00\1c\00\00\00=\02\00\00\0f\00\00\00\07\00\00\00\0c\00\00\00\04\00\00\00\08\00\00\00\03\00\00\00\08\00\00\00\04\00\00\00\09\00\00\00\0a\00\00\00\10\00\00\00\04\00\00\00\0b\00\00\00\0c\00\00\00\03\00\00\00\08\00\00\00\04\00\00\00\0d\00\00\00\0e\00\00\00\03\00\00\00\00\00\00\00\01\00\00\00\0f\00\00\00library/alloc/src/raw_vec.rscapacity overflow\00\00\00P\01\10\00\11\00\00\004\01\10\00\1c\00\00\00\06\02\00\00\05\00\00\00\11\00\00\00\00\00\00\00\01\00\00\00\12\00\00\0000010203040506070809101112131415161718192021222324252627282930313233343536373839404142434445464748495051525354555657585960616263646566676869707172737475767778798081828384858687888990919293949596979899\13\00\00\00\08\00\00\00\04\00\00\00\14\00\00\00\15\00\00\00\13\00\00\00\08\00\00\00\04\00\00\00\16\00\00\00Region pointer is null/usr/local/cargo/registry/src/github.com-1ecc6299db9ec823/cosmwasm-std-0.14.1/src/memory.rs\00\00\00\8e\02\10\00[\00\00\009\00\00\00\05\00\00\00Region starts at null pointer\00\00\00\8e\02\10\00[\00\00\00?\00\00\00\05"))
