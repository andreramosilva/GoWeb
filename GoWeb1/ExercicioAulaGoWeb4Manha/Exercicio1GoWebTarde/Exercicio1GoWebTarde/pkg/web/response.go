package web

// É necessário implementar um gerenciamento de resposta genérico para enviar sempre o mesmo formato nas solicitações. Para conseguir isso, as seguintes etapas devem ser executadas:

// Construa o webpack dentro do diretório pkg.
// Faça a estrutura Response com os capos: código, dados e erro:
// code terá o código de retorno.
// os dados terão a estrutura enviada pela aplicação (caso não haja erro).
// error terá o erro recebido em formato de texto (caso haja erro).
// Desenvolva uma função que receba o código como um inteiro, os dados como uma interface e o erro como uma string.
// A função deve retornar com base no código, se for uma resposta com os dados ou com o erro.
// Implemente esta função em todos os retornos dos controllers, antes de enviar a resposta ao cliente, a função deve gerar a estrutura que definimos.

type Response struct {
	codigo int         `json:"codigo"`
	dados  interface{} `json:"dados"`
	erro   string      `json:"erro"`
}

func newResponse(code int, data interface{}, err Response) {
	if code < 300 {
		return Response{
			codigo: code,
			dados:  data,
			erro:   "",
		}
	}

	return Response{
		codigo: code,
		dados:  "",
		erro:   err,
	}

}
