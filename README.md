# Homebroker

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


