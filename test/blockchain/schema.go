package blockchain

func MainRoute(router Router) {
	router.Get("/blockchain/", GetBlockChainArrayController)
	router.Get("/blockchain/hash/", GetBlockChainByHashController)
	router.Get("/blockchain/index/", GetBlockChainByIndexController)
	router.Get("/blockchain/validate", ValidateBlockChainController)
	router.Post("/blockchain/add", AddBlockChainController)
}
