package nex

import (
	"github.com/BDMCGaming/disney-infinity-1-nex/globals"
	commonnattraversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	commonsecure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	nattraversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"

	commondatastore "github.com/PretendoNetwork/nex-protocols-common-go/v2/datastore"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	commonSecureProtocol := commonsecure.NewCommonProtocol(secureProtocol)
	commonSecureProtocol.EnableInsecureRegister()

	dataStoreProtocol := datastore.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(dataStoreProtocol)
	commonDatastoreProtocol := commondatastore.NewCommonProtocol(dataStoreProtocol)
	commonDatastoreProtocol.SetManager(globals.DatastoreManager)

	natTraversalProtocol := nattraversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	commonnattraversal.NewCommonProtocol(natTraversalProtocol)
}
