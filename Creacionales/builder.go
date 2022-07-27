package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() BuildProcess
}

type ManufacturingDirector struct {
}

func (f *ManufacturingDirector) Construct() {
	//implementation
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	//implementation
}

type CarBuilder struct{}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}

func (c *CarBuilder) SetStructure() BuildProcess {
	return nil
}

func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}

type BikeBuilder struct{}

func (b *BikeBuilder) SetWheels() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetStructure() BuilldProcess {
	return nil
}

func (b *BikeBuillder) Build() VehicleProduct {
	return VehicleProduct{}
}
