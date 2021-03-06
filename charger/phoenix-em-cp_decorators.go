package charger

// Code generated by github.com/andig/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/andig/evcc/api"
)

func decoratePhoenixEMCP(base *PhoenixEMCP, meter func() (float64, error), meterEnergy func() (float64, error), meterCurrent func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterCurrent == nil && meterEnergy == nil:
		return base

	case meter != nil && meterCurrent == nil && meterEnergy == nil:
		return &struct {
			*PhoenixEMCP
			api.Meter
		}{
			PhoenixEMCP: base,
			Meter: &decoratePhoenixEMCPMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			*PhoenixEMCP
			api.MeterEnergy
		}{
			PhoenixEMCP: base,
			MeterEnergy: &decoratePhoenixEMCPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			*PhoenixEMCP
			api.Meter
			api.MeterEnergy
		}{
			PhoenixEMCP: base,
			Meter: &decoratePhoenixEMCPMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixEMCPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			*PhoenixEMCP
			api.MeterCurrent
		}{
			PhoenixEMCP: base,
			MeterCurrent: &decoratePhoenixEMCPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			*PhoenixEMCP
			api.Meter
			api.MeterCurrent
		}{
			PhoenixEMCP: base,
			Meter: &decoratePhoenixEMCPMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratePhoenixEMCPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			*PhoenixEMCP
			api.MeterCurrent
			api.MeterEnergy
		}{
			PhoenixEMCP: base,
			MeterCurrent: &decoratePhoenixEMCPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratePhoenixEMCPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			*PhoenixEMCP
			api.Meter
			api.MeterCurrent
			api.MeterEnergy
		}{
			PhoenixEMCP: base,
			Meter: &decoratePhoenixEMCPMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratePhoenixEMCPMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratePhoenixEMCPMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decoratePhoenixEMCPMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePhoenixEMCPMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePhoenixEMCPMeterCurrentImpl struct {
	meterCurrent func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEMCPMeterCurrentImpl) Currents() (float64, float64, float64, error) {
	return impl.meterCurrent()
}

type decoratePhoenixEMCPMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePhoenixEMCPMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}
