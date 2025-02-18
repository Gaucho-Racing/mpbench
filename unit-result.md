# Summary

| ID | Name | Status | Progress |
|----|------|--------|----------|
0x003 | ECU Status One Test 1 | ✅ PASS | 21/21 (100%)
0x003 | ECU Status One Test 2 | ✅ PASS | 15/15 (100%)
0x004 | ECU Status Two Test 1 | ❌ FAIL | 5/8 (63%)

# Failing Tests

### `0x004` ECU Status Two Test 1

**5/8** signals ingested correctly

Input Data: `{0x23, 0x00, 0x12, 0xf2, 0x2a, 0x19, 0x00, 0x81}`

| Signal Name | Decoded Value | Expected Value | Status |
|------------|---------------|----------------|---------|
| `ecu_temp` | 35.0 | 35.0 | ✅ |
| `ecu_voltage` | 12.2 | 14.4 | ❌ |
| `ecu_status` | 2 | 2 | ✅ |
| `ecu_error_code` | 242 | 0 | ❌ |
| `ecu_fuel_pressure` | 42 | 42 | ✅ |
| `ecu_oil_pressure` | 25 | 45 | ❌ |
| `ecu_throttle_pos` | 0 | 0 | ✅ |
| `ecu_engine_speed` | 129 | 129 | ✅ |
