[![](https://github.com/hilli/apcupsd_json_status/workflows/Build/badge.svg)](https://github.com/hilli/apcupsd_json_status/actions)

# apcupsd_json_status

The program is intended to talk to a `apcupsd` daemon, listing on it's NIS interface.
It will output the found parameters in a single line of JSON, suitable for many modern monitoring tools. My own usecase is for [Humio](https://www.humio.com). If you are using Prometheus, check out [Matt Layher](https://github.com/mdlayher)s [apcupsd_exporter](https://github.com/mdlayher/apcupsd_exporter) (He also created the library this project is using to collect the data from the apcupsd, so thanks Matt).

## Examples:

### Get a single status from the UPS

```
# ./apcupsd_json_status
{"timestamp":"2019-12-02T08:36:49.063416+01:00","USBName":"UPS1","UPSMode":"Stand Alone","UPSModel":"Smart-UPS 3000 RM","LoadPercent":33.1,"BatteryChargePercent":100,"LineVolts":231.8,"NominalInputVoltage":0,"BatteryVoltage":54,"NominalBatteryVoltage":48,"BatteryNumberTransfersTotal":0,"BatteryTimeLeftSeconds":1440,"BatteryTimeOnSeconds":0,"LastTransferOnBattery":"0001-01-01T00:00:00Z","LastTransferOffBattery":"0001-01-01T00:00:00Z","LastSelftest":"0001-01-01T00:00:00Z","NominalPowerWatts":0}
```

### Run as a daemon

```
# ./apcupsd_json_status -d
{"timestamp":"2019-12-02T08:36:49.063416+01:00","USBName":"UPS1","UPSMode":"Stand Alone","UPSModel":"Smart-UPS 3000 RM","LoadPercent":33.1,"BatteryChargePercent":100,"LineVolts":231.8,"NominalInputVoltage":0,"BatteryVoltage":54,"NominalBatteryVoltage":48,"BatteryNumberTransfersTotal":0,"BatteryTimeLeftSeconds":1440,"BatteryTimeOnSeconds":0,"LastTransferOnBattery":"0001-01-01T00:00:00Z","LastTransferOffBattery":"0001-01-01T00:00:00Z","LastSelftest":"0001-01-01T00:00:00Z","NominalPowerWatts":0}
{"timestamp":"2019-12-02T08:37:09.184618+01:00","USBName":"UPS1","UPSMode":"Stand Alone","UPSModel":"Smart-UPS 3000 RM","LoadPercent":33.8,"BatteryChargePercent":100,"LineVolts":231.8,"NominalInputVoltage":0,"BatteryVoltage":51.6,"NominalBatteryVoltage":48,"BatteryNumberTransfersTotal":0,"BatteryTimeLeftSeconds":1500,"BatteryTimeOnSeconds":0,"LastTransferOnBattery":"0001-01-01T00:00:00Z","LastTransferOffBattery":"0001-01-01T00:00:00Z","LastSelftest":"0001-01-01T00:00:00Z","NominalPowerWatts":0}
etc...
```

### Specific connection to the apcupsd

```
# ./apcupsd_json_status -d -c ups-host:3351
```

Default it will try to connect to localhost:3351.