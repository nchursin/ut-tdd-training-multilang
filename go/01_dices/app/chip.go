package app

type Chip struct{ Amount int }

func (c Chip) Equal(other Chip) bool          { return c.Amount == other.Amount }
func (c Chip) GreaterOrEqual(other Chip) bool { return c.Amount >= other.Amount }
func (c Chip) LessOrEqual(other Chip) bool    { return c.Amount <= other.Amount }
func (c Chip) Add(other Chip) Chip            { return Chip{Amount: c.Amount + other.Amount} }
func (c Chip) Subtract(other Chip) Chip       { return Chip{Amount: c.Amount - other.Amount} }
func (c Chip) Multiply(value int) Chip        { return Chip{Amount: c.Amount * value} }
