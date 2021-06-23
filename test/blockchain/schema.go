package blockchain

func MainRoute(router Router) {
	router.Get("/blockchain/", GetBlockChainArrayController)
	router.Post("/blockchain/add", AddBlockChainController)
	router.Get("/blockchain/hash/", GetBlockChainByHashController)
	router.Get("/blockchain/index/", GetBlockChainByIndexController)
	router.Post("/blockchain/validate", ValidateBlockChainController)
}
