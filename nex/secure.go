package nex

import (
	"fmt"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	"os"
	"strconv"

	"encoding/hex"
	"github.com/BDMCGaming/disney-infinity-1-nex/globals"
	"github.com/PretendoNetwork/nex-go/v2"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()
	globals.SecureServer.ByteStreamSettings.UseStructureHeader = true

	globals.SecureEndpoint = nex.NewPRUDPEndPoint(1)
	globals.SecureEndpoint.IsSecureEndPoint = true
	globals.SecureEndpoint.ServerAccount = globals.SecureServerAccount
	globals.SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)

	globals.SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 5, 2))
	globals.SecureServer.AccessKey = "46f82708"

	globals.SecureEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("==Disney Infinity 1.0 - Secure==")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID)
		fmt.Printf("Method ID: %d\n", request.MethodID)
		fmt.Printf("Data: %s\n", hex.EncodeToString(request.Parameters))
		fmt.Println("===============")
	})

	globals.SecureEndpoint.OnError(func(err *nex.Error) {
		globals.Logger.Errorf("Secure: %v", err)
	})

	globals.DatastoreManager = common_globals.NewDataStoreManager(globals.SecureEndpoint, globals.Postgres)
	globals.DatastoreManager.SetS3Config(globals.S3Bucket, globals.S3KeyBase, globals.S3Manager)
	globals.DatastoreManager.GetUserFriendPIDs = globals.GetUserFriendPIDs

	registerCommonSecureServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_INFINITY1_SECURE_SERVER_PORT"))

	globals.SecureServer.Listen(port)
}
