package frequencias

/*func NovaFrequencia(c *gin.Context) {
	fmt.Println("Tentando cadastrar uma nova frequencia")
	req := modelApresentacao.Frequencia{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create. Parameters were not passed correctly",
			"err":     err.Error(),
		})
		return
	}

	err := frequencias.NovaFrequencia(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Frequencia adicionada com sucesso!"})
}
*/

/*func ListarFrequenciasUsuario(c *gin.Context) {
	fmt.Println("Tentando buscar as frequencias de um usuário")
	idUser, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"err": err.Error()})
		return
	}

	frequencias, err := frequencias.ListarFrequenciasUsuario(utils.GetInt64Pointer(int64(idUser)))
	if err != nil {
		c.JSON(404, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, frequencias)
}
*/