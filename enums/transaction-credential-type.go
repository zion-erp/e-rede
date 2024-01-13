package enums

type TransactionCredentialType string

const (
	TransactionCredentialType_CardOnFile                          TransactionCredentialType = "01" //Card on File (Credencial armazenada)	O consumidor concorda que seu cartão seja armazenado com o comerciante para futuras transações que possam ocorrer de tempos em tempos.	Transações de aplicativos de carro.
	TransactionCredentialType_RecurrentPaymentCustomerAgreement   TransactionCredentialType = "02" //Ordem Permanente	O consumidor concorda que seu cartão seja armazenado com o comerciante e inicia uma primeira transação em uma série destinada a um valor variável e uma frequência fixa.	Pagamento mensal de serviços.
	TransactionCredentialType_SubscriptionCustomerAgreement       TransactionCredentialType = "03" //Assinatura	O consumidor concorda que seu cartão seja armazenado e inicia uma primeira transação em uma série destinada a um valor fixo e uma frequência fixa.	Assinatura mensal de jornal.
	TransactionCredentialType_InstallmentCustomerAgreement        TransactionCredentialType = "04" //Parcelado	O consumidor concorda que seu cartão seja armazenado para estabelecer um plano de parcelamento e inicia uma primeira transação em uma série.	A transações parceladas.
	TransactionCredentialType_RecurrentPaymentNoCustomerAgreement TransactionCredentialType = "06" //Ordem Permanente	Usar os dados da conta do titular do cartão para uma transação que deve ocorrer em intervalos regulares por um valor variável.	Pagamentos mensais de serviços
	TransactionCredentialType_SubscriptionNoCustomerAgreement     TransactionCredentialType = "07" //Assinatura	Usar os dados da conta do titular do cartão para uma transação que deve ocorrer em intervalos regulares por um valor fixo.	Assinatura mensal ou pagamento de serviço mensal fixo.
	TransactionCredentialType_InstallmentNoCustomerAgreement      TransactionCredentialType = "08" //Parcelado	Armazenar os dados da conta do titular do cartão para uso do comerciante para iniciar uma ou mais transações futuras por um valor conhecido com uma determinada duração com base em uma única compra.	Comprar uma TV por R$ 1.000, pagando em quatro parcelas iguais de R$ 250 (a primeira transação é CIT, as três transações restantes são MIT).
	TransactionCredentialType_LateCharge                          TransactionCredentialType = "10" //Cobrança atrasada	Uma cobrança adicional da conta após a prestação dos serviços iniciais e o processamento do pagamento.	Cobrança do frigobar do hotel após o titular do cartão fazer check out do hotel.
	TransactionCredentialType_Fine                                TransactionCredentialType = "11" //No show (Multa)	Uma multa cobrada de acordo com a política de cancelamento do comerciante.	O cancelamento de uma reserva pelo titular do cartão sem aviso prévio adequado ao comerciante.
	TransactionCredentialType_Resend                              TransactionCredentialType = "12" //Reenvio	A tentativa de obter autorização para uma transação que foi recusada, mas a resposta do emissor não proíbe que o comerciante tente mais tarde.	Fundos insuficientes/ resposta acima do limite de crédito/ retentativa de transações de trânsito.
)
