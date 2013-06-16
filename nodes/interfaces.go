package nodes

type NodeInterface interface {
  Left() interface{}
  Right() interface{}
}

type ComparisonInterface interface {
  NodeInterface
  Attribute() AttributeInterface
  Value() interface{}
  Or(ComparisonInterface) OrInterface
}

type ComparableInterface interface {
  NodeInterface
  Eq(interface{}) ComparisonInterface
  Neq(interface{}) ComparisonInterface
  Gt(interface{}) ComparisonInterface
  Gte(interface{}) ComparisonInterface
  Lt(interface{}) ComparisonInterface
  Lte(interface{}) ComparisonInterface
  Matches(interface{}) ComparisonInterface
}

type OrInterface interface {
  NodeInterface
}

type ReferenceInterface interface {
  NodeInterface
  Name() string
  Aliases() []string
  AddAliases(...string) ReferenceInterface
}

type AttributeInterface interface {
  ComparableInterface
  Name() string
  Reference() ReferenceInterface
  TableName() string
}