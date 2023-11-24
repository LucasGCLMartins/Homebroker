# Homebroker

[![Watch the video](https://i.imgur.com/905z5C5.png)](https://www.youtube.com/watch?v=ys00FOJ9Eno)

# COMO RODAR
  -É necessário executar 3 contêiners para cada uma das pastas do programa e elas devem ser executadas na ordem abaixo:<br>
    - GO<br>
    - Nest<br>
    - Next<br>

# Semana 25/09 até 29/09
  -Ambiente teste do Go está funcionando e a API foi testada com sucesso, ao pesquisar um ativo por nome ele retorna o historico de preços dos ultimos 30 dias registrados<br>Exemplo:<br>
  
![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/4ef9e683-9a75-4c43-aaab-d448c41e91e9)

  -Apache Kafka configurado na maquina (Utilizando Windows)<br>
  Kafka é rodado em duas partes, o Zookeper que sincroniza os clusteres e o Broker que gerencia os topicos
  
  Zookeper Kafka rodando:<br>
  C:\kafka_2.13-3.5.0>.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties
  ![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/2a87d5e0-b783-420e-b483-57b37f9f4353)

  Broker Kafka rodando:<br>
  C:\kafka_2.13-3.5.0>.\bin\windows\kafka-server-start.bat .\config\server.properties
  ![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/6cd07bba-0c4e-4d57-951e-8ed8457af100)

Topico criado:
![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/439d19f8-7618-4762-9f7a-f18b4f9c45f6)

Produtor Criado:
![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/110dddcf-0f38-4a94-9f73-95d90b64c48d)

Consumidor criado
![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/b794140a-9b9f-41d0-896d-1f63a9dfbbc8)

# Semana 16/09 até 20/09

  -Nessa semana foi criado os testes das operações de compra e venda das ações. Esses testes foram automatizados utilizando testify, que é um framework para testes da propria linguagem go.
  Para realização dos testes foram escritos 5 casos diferentes:
Caso de Teste | Resultado |
-- | -- | 
Compra de ativo | OK
Pedido de compra que não tem venda | OK
Compra parcial | OK
Compra com dois preços | OK
Não foi possivel fazer a compra | OK


![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/97a79af3-2ce0-40f1-88f8-e3f10905cbb3)

Além disso também foi realizada a integração do codigo com o Apache Kafka, demonstrado em um vídeo

https://github.com/LucasGCLMartins/Homebroker/assets/73212163/ec5c84c5-203e-4652-aa99-262209eb80e8

Nesse video, primeiro eu mostro que o Kafka está rodando no docker
![image](https://github.com/LucasGCLMartins/Homebroker/assets/73212163/af669f94-9cb4-4a9e-9d7d-86bd39cf4279)

Em seguida eu faço uma simulação das operações que podem ser efetuadas na plataforma, na pagina do localhost eu crio duas ordens, uma de compra e outra de venda, e mostro que assim que elas foram concluidas o programa registrou e printou no terminal.
Como as duas ordens completam uma a outra, o algoritmo de venda resolveu que eram compativeis e concluiu a transação, assim registrando os status dela como CLOSED

