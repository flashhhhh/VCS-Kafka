# Kafka

##  Khái niệm
Apache Kafka là một nền tảng phân tán xử lý luồng dữ liệu cho việc nhập dữ liệu và xử lý dữ liệu thời gian thực. Streaming data là dữ liệu được sinh ra liên tục từ hàng nghìn nguồn dữ liệu, và dữ liệu thường được gửi đi ngay lập tức. Một nền tảng streamming cần xử lý với khối lượng dữ liệu khổng lồ này, và xử lý chúng tuần tự và liên tục.

Kafka cung cấp 3 chức năng chính cho người dùng:
+ Publish và subcribe tới luồng dữ liệu
+ Lưu trữ luồng dữ liệu hiệu quả theo thứ tự mà dữ liệu được sinh ra
+ Xử lý các luồng dữ liệu trong thời gian thực

## Mục đích sử dụng
### Gửi và nhận tin nhắn
Kafka có thể thay thế hiệu quả cho các message broker truyền thống khác. Message brokers được sử dụng trong nhiều trường hợp (tách biệt việc xử lý dữ liệu khỏi nguồn dữ liệu, làm bộ đệm lưu trữ các tin nhắn chưa được xử lý, ...).

So với hầu hết các hệ thống messaging, Kafka có thông lượng tốt hơn, có các cơ chế như các partition, replication giúp khả năng xử lý với mất mát dữ liệu tốt hơn, giúp Kafka trở thành giải pháp hiệu quả cho các ứng dụng yêu cầu xử lý thông lượng tin nhắn lớn.

### Metrics
Kafka có thể được dùng để giám sát hoạt động của ứng dụng từ dữ liệu của ứng dụng đó. Điều này bao gồm tổng hợp các thống kê từ các ứng dụng bị phân tán ở nhiều chỗ về một nguồn cấp dữ liệu hoạt động tập trung.

### Tổng hợp log
Tổng hợp log thường thu thập các log files từ servers và lưu trữ ở một nơi tập trung để xử lý. Kafka tóm tắt các chi tiết của các tệp và cung cấp một bản tóm tắt gọn hơn về dữ liệu nhật ký hoặc sự kiện dưới dạng một luồng tin nhắn, cho phép xử lý dữ liệu ít độ trễ hơn và hỗ trợ tốt hơn cho nhiều nguồn dữ liệu và nhiều nguồn tiêu thụ phi tập trung.

So với các hệ thống ghi log tập trung như Scribe hoặc Flume, Kafka cung cấp hiệu suất tốt như nhau, đảm bảo dữ liệu không bị mất do khả năng replication và độ trễ đầu cuối thấp hơn nhiều.

### Stream Processing
So với việc xử lý dữ liệu theo batch yêu cầu dữ liệu đầu vào phải là một khối, khiến cho việc xử lý dữ liệu không được liên tục, Kafka cho phép dữ liệu mới được sinh ra được truyền trực tiếp đến nơi xử lý dữ liệu, giúp việc nhập hay xử lý dữ liệu được diễn ra trong thời gian thực

### Kiến trúc Event-Driven
Dùng Kafka làm trung gian để các dịch vụ có thể giao tiếp mà không phụ thuộc trực tiếp vào nhau.

## So sánh Kafka với các Message Queue khác
### Kafka vs RabbitMQ
| Tính chất    | Apache Kafka | RabbitMQ |
| -------- | ------- | ------ |
| Kiến trúc | **Kafka sử dụng mô hình nhật ký phân vùng, nên có thể kết hợp theo cả 2 hướng message queue và publish subcribe** | RabbitMQ sử dụng message queue |
| Khả năng mở rộng | **Cho phép các partition được phân tán trên nhiều server khác nhau** | Tăng số lượng consumer để giảm bớt công việc cho các consumer khác |
| Thời gian lưu trữ tin nhắn | **Theo quy tắc được định nghĩa, ví dụ như 1 ngày, hoặc theo config của người dùng** | Theo ack, tin nhắn sẽ bị xóa ngay lập tức khi có 1 consumer tiêu thụ thành công |
| Nhiều consumer | **Nhiều consumer có thể subcribe tới cùng 1 topic, do Kafka cho phép cùng 1 tin nhắn có thể được gửi lại cho các consumer khác trong 1 khoảng thời gian** | Nhiều consumer không thể tiêu thụ cùng 1 tin nhắn, do tin nhắn được xóa ngay lập tức khi nó được tiêu thụ |
| Replication | Các topic được mặc định là nhân bản, nhưng người dùng có thể setup để không nhân bản cho topic đó | Tin nhắn được mặc định là không nhân bản, nhưng người dùng có thể setup để nhân bản |
| Thứ tự tin nhắn | **Các consumer nhận được tin nhắn theo đúng thứ tự đầu vào do kiến trúc nhật ký phân vùng** | Tin nhắn được chuyển đến cho consumer theo thứ tự đến hàng đợi. Nếu có nhiều consumer, mỗi consumer sẽ xử lý một tập con của tin nhắn trong hàng đợi |
| Giao thức | Kafka sử dụng giao thức nhị phân trên TCP | **Giao thức hàng đợi nhắn tin nâng cao (AMQP) với hỗ trợ thông qua các plugin: MQTT, Stomp**|

### Kafka vs Redis
| Tính chất    | Apache Kafka | RabbitMQ |
| -------- | ------- | ------ |
| Kích thước tin nhắn | **Hỗ trợ tin nhắn với kích thước lên đến 1GB với khả năng nén và lưu trữ** | Hỗ trợ tin nhắn với dung lượng nhỏ hơn |
| Cách nhận tin nhắn | Subcriber pull tin nhắn về từ hàng đợi | Publisher push tin nhắn đến subscriber |
| Khả năng lưu tin nhắn | **Lưu tin nhắn sau khi được truy xuất** | Không lưu lại tin nhắn |
| Khả năng xử lý lỗi | **Có khả năng xử lý lỗi mạnh mẽ ở tầng gửi nhận tin nhắn.** | Người dùng phải xử lý lỗi ở tầng ứng dụng (bằng timeout, giới hạn client và kích thước bộ nhớ đệm) |
| Khả năng chạy song song | **Kafka hỗ trợ song song hóa khi nhiều consumer có thể truy cập đồng thời cùng 1 tin nhắn** | Không hỗ trợ song song hóa |
| Thông lượng | **Có thông lượng cao hơn do khả năng đọc/ ghi bất đồng bộ** | Chậm hơn do Redis phải đợi phản hồi thành công trước khi gửi đi tin nhắn tiếp theo |
| Độ trễ | Có độ trễ thấp, tuy nhiên chậm hơn Redis do phải replicate dữ liệu | **Độ trễ cực thấp, đặc biệt khi gửi những tin nhắn với kích thước bé** |
| Khả năng chịu lỗi | **Có bản sao của dữ liệu thông qua cơ chế partition** | Không hỗ trợ backup |

## Kafka components
### 1. Kafka Producer
Kafka producer là ứng dụng client dùng để publish tin nhắn đến Kafka topics. Kafka sẽ gán partition và offset cho tin nhắn đó khi nhận được tin nhắn.

#### Message Key và Partitioning
Do Kafka tăng khả năng scale và hiệu năng bằng cách viết tin nhắn đồng thời vào nhiều partition, tuy nhiên điều này làm mất thứ tự của tin nhắn. Thứ tự tin nhắn chỉ được đảm bảo đúng thứ tự trong cùng một partition. Do đó, ta có thể thêm trường **message key** vào tin nhắn. Các tin nhắn có cùng **message key** được đảm bảo sẽ vào cùng 1 partition.

#### Durability và Error Handling
Kafka sử dụng **acks** để kiểm soát có bao nhiêu brokers phải phản hồi về khi được producer gửi tin nhắn.

Với **acks = all**, toàn bộ replicas của partition đó phải nhận được tin nhắn và phản hồi về, khi đó producer mới công nhận tin nhắn được gửi thành công. Nếu đặt **acks = 0**, producer sẽ tiếp tục gửi tin nhắn mà không quan tâm đến việc có replica nào nhận được tin nhắn chưa.

Có thể setup thêm một số options cho timeout:
* **retries**: Producer sẽ gửi lại tối đa bao nhiêu lần khi tin nhắn bị gửi lỗi.
* **retry.backoff.ms**: Producer sẽ đợi bao lâu trước khi bắt đầu gửi lại.
* **delivery.timeout.ms**: Producer sẽ đợi một acknowledgment trong bao lâu.

### 2. Kafka Consumer Group
Các Kafka Consumers có thể hoạt động thành 1 Consumer Group khi cùng consume 1 topic.

Kafka mạnh về scale, do đó có thể có nhiều producers cùng viết vào topic và nhiều consumers cùng tiêu thụ topic. Tuy nhiên, nếu nhiều producers gửi quá nhiều tin nhắn đến 1 topic, consumer không thể xử lý kịp. Do đó, cần consumer group để cho phép nhiều consumers cùng tiêu thụ cùng 1 topic mà chia đều tin nhắn cho các consumers đó, đảm bảo rằng mỗi tin nhắn chỉ có 1 consumer xử lý.

#### Cách hoạt động
Với mỗi partition chỉ được tiêu thụ bởi đúng 1 consumer trong group. Do đó, mỗi consumer tiêu thụ 1 tập con các partition của topic đó (Nếu như số lượng partitions không ít hơn số lượng consumers)

Khi có nhiều consumers hơn partitions, một vài consumer sẽ không làm gì cả. Tuy nhiên nếu như có 1 consumer bị fail, consumer rảnh đó sẽ ngay lập tức thế chỗ.

Khi có nhiều consumer groups, Kafka đảm bảo các groups đó có thể consume cùng một lượng tin nhắn.

#### Cách phân chia partitions trong 1 Consumer Group
**Group Coordinator** giúp cho việc phân phối partitions giữa các consumer trong 1 group đồng đều nhất có thể, kể cả khi có thay đổi. 

**Group Coordinator** thực hiện các chức năng sau:
* Khi có 1 consumer tham gia hoặc rời khỏi group, group coordinator phân chia lại các partition dựa trên thuật toán cân bằng.
* Các consumer định kỳ sẽ gửi về 1 **heartbeat**. Nếu trong một khoảng timeout mà không nhận được **heartbeat** của 1 consumer, **Group Coordinator** sẽ coi nó là không hoạt động và bắt đầu tái cân bằng lại partitions trong group.
* **Group Coordinator** theo dõi các tin nhắn mà **consumer group** đã tiêu thụ bằng cách lưu các offsets của từng consumer vào 1 topic tên **__consumer_offsets**. Khi một consumer khởi động lại, nó sẽ yêu cầu coordinator gửi lại offset cuối cùng của nó, tránh việc xử lý một tin nhắn nhiều lần.

Các chiến thuật phân công partitions cho các consumers của Kafka:
* **Range**: Mỗi consumer được cấp các partitions liên tiếp nhau của topic. Điều này có thể làm mất cân bằng giữa các consumer.
Ví dụ: Có 3 consumers trong 1 group là C1, C2 và C3, subscribe đến 3 topics là T1, T2 và T3. Nếu T1 có 3 partitions, T2 và T3 cùng có 2 partitions.
	* C1 sẽ tiêu thụ:
		* Partition thứ nhất của T1
		* Partition thứ nhất của T2
		* Partition thứ nhất của T3
	* C2 sẽ tiêu thụ:
		* Partition thứ hai của T1
		* Partition thứ hai của T2
	* C3 sẽ tiêu thụ:
		* Partition thứ ba của T1
		* Partition thứ ba của T2

* **Round-Robin**: Tất cả các partitions từ tất cả các topic được subsribe sẽ được cấp cho các consumers 1 cách tuần tự. Điều này giúp giảm sự mất cân bằng giữa các partition hơn.
Ví dụ: Có 3 consumers trong 1 group là C1, C2 và C3, subscribe đến 3 topics là T1, T2 và T3. Nếu T1 có 3 partitions, T2 và T3 cùng có 2 partitions.
	* C1 sẽ tiêu thụ:
		* Partition thứ nhất của T1
		* Partition thứ nhất của T2
		* Partition thứ hai của T3
	* C2 sẽ tiêu thụ:
		* Partition thứ hai của T1
		* Partition thứ hai của T2
	* C3 sẽ tiêu thụ:
		* Partition thứ ba của T1
		* Partition thứ nhất của T3

### 3. Kafka Broker
Broker là node lõi tính toán chính của Kafka. Producers gửi tin nhắn đến **Kafka Brokers**, và broker lưu trữ dữ liệu trên nhiều partitions. Consumers sẽ kết nối đến các brokers này để fetch dữ liệu từ các topics được đăng ký.

#### Các chức năng chính của **Kafka broker**
##### Message Management
Khi producer gửi 1 tin nhắn đến 1 topic, broker sẽ quyết định lưu tin nhắn ở partition nào thông qua message key của tin nhắn đó, hoặc thông qua round robin.

Ngoài ra, brokers còn theo dõi offsets cho mỗi partition, để biết rằng tin nhắn nào đã được consume và tin nhắn nào chưa. Điều này giúp consumers có thể tiếp tục tiêu thụ tin nhắn kể cả sau khi bị mất kết nối.

##### Replication
Kafka có thể nhân bản mỗi partition trên nhiều brokers, và một trong số các broker được chọn làm leader của riêng partition đó. Leader chịu trách nhiệm cho toàn bộ yêu cầu đọc và viết, trong khi các followers phải đồng bộ dữ liệu của nó theo leader.

Khi leader sập, một trong số các follower sẽ được chọn làm leader mới.

#### Quản lý Kafka Broker trong 1 cụm
Kafka sử dụng Zookeeper (Cũ) / KRaft (Mới) để kiểm soát metadata cũng như trạng thái của các brokers.

##### Metadata log
Toàn bộ thay đổi về metadata trong 1 cụm sẽ được lưu trữ như 1 chuỗi sự kiện trong 1 Kafka topic.

##### Leader Election
Một trong số các broker được bầu ra làm leader, với nhiệm vụ thực thi các yêu cầu ghi và nhân bản các thay đổi đến cho brokers khác, và ghi vào metadata log.

Định kỳ, metadata log được leader snapshot lại để tránh việc tốn quá nhiều bộ nhớ, đồng thời để khôi phục nhanh hơn khi có sự cố xảy ra.

### 4. Kafka Topic
Kafka topic là 1 append-only file để lưu trữ tin nhắn của producer gửi vào. Các producers sẽ ghi tin nhắn vào topic, và consumer group sẽ pull tin nhắn về.

Topic partitions là đơn vị lưu trữ tin nhắn cơ bản nhất của Kafka. Trong mỗi partition, Kafka chia nhỏ dữ liệu thành các **segment**. Mỗi segment hoạt động giống như 1 file. Chỉ có 1 segment hoạt động trong 1 khoảng thời gian, và kafka ghi dữ liệu mới nhất vào segment đó. Có thể kiểm soát **segment** thông qua config:
* **log.segment.bytes**: kiểm soát dung lượng lớn nhất trước khi commit 1 segment.
* **log.segment.ms**: kiểm soát thời gian tối đa Kafka mở 1 segment.

Ngoài ra, topic còn hỗ trợ việc nén dữ liệu để tối ưu bộ nhớ. Có thể nén dữ liệu mặc định của tầng topic hoặc nén dữ liệu thông qua kiểu nén dữ liệu của producer truyền vào. Có thể thay đổi kiểu nén dữ liệu mặc định thông qua config **compression.type**.

### 5. Kafka Partition
Kafka cho phép chọn cách producers publish tin nhắn đến partition, và cách consumers tiêu thụ partitions.

#### Chiến thuật partition cho producer
##### Default
Khi **message key** là null, Kafka lựa chọn ngẫu nhiên 1 partition của topic đó để gửi tin nhắn.

Khi **message key** tồn tại, Kafka map tin nhắn đến partition giống với mã hóa hash của **message key** đó. Tuy nhiên, điều này chỉ đúng khi số lượng partition không đổi. Khi thêm partition mới vào, Kafka sẽ viết đến partition khác.

##### Round-Robin
Kafka ghi tin nhắn lần lượt vào các topic xong lại quay vòng. Tin nhắn cùng **message key** vẫn được đảm bảo gửi đến cùng partition.

##### Custom
Ta có thể truyền hàm do ta tự viết tùy thuộc vào ngữ cảnh.

### 6. Kafka Offset
Kafka sử dụng offset để theo dõi tin nhắn từ lúc bắt đầu ghi đến khi xử lý hết các tin nhắn, giúp cho các consumer có thể lấy đúng tin nhắn theo đúng thứ tự được producer gửi.

Trong khi tin nhắn nào cũng có offset, Kafka lưu một vài offset đặc biệt:
* **Log-end offset**: Tin nhắn cuối cùng xuất hiện trong 1 partition.
* **High watermark offset**: Vị trí cuối cùng mà các tin nhắn đã được đồng bộ thành công đến các replica.
* **Committed offset**: Tin nhắn cuối cùng được 1 consumer xử lý thành công.

Kafka có thể dùng offsets cho nhiều mục đích khác nhau.
#### Quản lý replication của từng partition
Khi tin nhắn gửi đến topic, nó sẽ được nhân bản ra các partition khác. Khi consumer fetch tin nhắn, nó sẽ chỉ quan tâm đến các tin nhắn trước **high watermark offset**, tức là các tin nhắn đã được đồng bộ thành công đến toàn bộ replica.
#### Quản lý lỗi của consumer
Offset giúp consumer bị lỗi và khởi động lại có thể biết tin nhắn tiếp theo phải xử lý. Có thể tắt auto commit, có nghĩa là chỉ khi consumer commit là dã xử lý thành công thì Kafka mới tăng offset.
#### Đảm bảo chuyển phát tin nhắn với Kafka offset
Kafka hỗ trợ 3 mức độ tin cậy.
##### At Most Once
Consumer sẽ auto commit tin nhắn ngay khi nó nhận được tin nhắn, trước khi xử lý tin nhắn đó.
Do đó, nếu tin nhắn bị xử lý lỗi, tin nhắn đó sẽ bị mất. Tuy nhiên, điều này đảm bảo mỗi tin nhắn chỉ được gửi đi đúng một lần duy nhất.
##### At Least Once
Consumer chỉ commit nhận được tin nhắn sau khi xử lý xong tin nhắn đó. Nếu như consumer xử lý tin nhắn thành công nhưng gửi yêu cầu commit bị lỗi, Kafka sẽ gửi lại tin nhắn đó đến consumer.
Do đó, tin nhắn có thể bị gửi lặp lại, nhưng đảm bảo được không bị mất tin nhắn nào.
##### Exactly Once
Kafka sẽ tạo 1 transaction gồm việc xử lý tin nhắn và việc gửi commit. Khi bị lỗi ở bước nào, Kafka sẽ thực hiện lại từ đầu.
Do đó, đảm bảo được mỗi 1 tin nhắn sẽ được xử lý đúng 1 lần và không bị mất tin nhắn. Tuy nhiên, điều này ảnh hưởng đến hiệu năng và phức tạp trong việc triển khai.
#### Giám sát consumer lag
Consumer lag là khoảng cách giữa **committed_offset** và **log-end offset**. Nếu lag này quá lớn có thể phá vỡ hệ thống. Điều này thường xảy ra khi:
* Phân chia không đồng đều các partitions cho các consumers trong 1 group.
* Tốc độ gửi của producer nhanh hơn nhiều so với tốc độ xử lý của consumer.