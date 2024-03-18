package aircraftlib

import (
	"github.com/homier/go-capnp/v3"
)

var (
	// Make sure interface types satisfy TypeParam:
	_ capnp.TypeParam[Echo] = Echo{}

	// ...and structs:
	_ capnp.TypeParam[Zdate] = Zdate{}

	// ...and lists:
	_ capnp.TypeParam[Echo_List]  = Echo_List{}
	_ capnp.TypeParam[Zdate_List] = Zdate_List{}
)
