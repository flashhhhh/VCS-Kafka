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

## So sánh điểm mạnh, điểm yếu với các Message Queue khác
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
### 1. Kafka Broker
Một broker là một server của Kafka để lưu trữ dữ liệu và phục vụ các request từ client.

* Một cụm Kafka bao gồm nhiều brokers để tăng khả năng chịu lỗi và khả năng mở rộng.
* Mỗi broker lưu các partition của nhiều topic khác nhau.
* Các Kafka brokers hoạt động với nhau để cân bằng tải và đảm bảo khả năng chịu lỗi tốt.

### **2. Topic**
- Danh mục logic để phân loại tin nhắn
- Được chia thành các **partition** để xử lý song song

### **3. Partition**
- Chuỗi tin nhắn có thứ tự, bất biến
- Mỗi tin nhắn có **offset** riêng (ID tuần tự)
- Được sao chép trên nhiều broker (leader + follower)

### **4. Producer**
- Ứng dụng gửi tin nhắn đến topic
- Khi gửi tin nhắn đi thì tin nhắn có thể ở 1 partition bất kỳ của topic thông qua Round-Robin, phân loại qua key hoặc theo setup của người dùng.
- Cấu hình **acks** (0, 1 hoặc `all`) cho độ tin cậy:
	* acks = 0: gửi đi mà không cần phản hồi từ brokers.
	* acks = 1: gửi đi và chờ ít nhất 1 broker phản hồi ack (nhận được tin nhắn).
	* acks = all: gửi đi và chờ đến khi tất cả brokers phản hồi ack (nhận được tin nhắn).

### **5. Consumer**
- Đọc tin nhắn từ topic, mỗi consumer tiêu thụ 1 partition của 1 topic.
- Theo dõi tiến trình qua **offset**, mỗi consumer có 1 chỉ số **offset** riêng biệt.
- Tổ chức thành **consumer group** để xử lý song song, tức là 1 tin nhắn chỉ được tiêu thụ bởi đúng 1 consumer trong 1 consumer group.

### **6. Consumer Group**
- Nhóm consumer chia sẻ tải
- Mỗi partition chỉ được xử lý bởi **một** consumer trong nhóm
- Tự động cân bằng lại khi có lỗi

### **7. ZooKeeper (Cũ) / KRaft (Mới)**
- **ZooKeeper**: Quản lý metadata, bầu chọn leader (trước Kafka 3.0)
- **KRaft**: Giao thức đồng thuận tích hợp (Kafka 3.0+), không cần ZooKeeper

### **8. Kafka Connect**
- Khung tích hợp với hệ thống bên ngoài
  - **Source Connector**: Nhận dữ liệu (VD: DB → Kafka)
  - **Sink Connector**: Xuất dữ liệu (VD: Kafka → S3)

### **9. Kafka Streams**
- Thư viện xử lý luồng thời gian thực
- Hỗ trợ các thao tác như `lọc`, `tổng hợp`, `kết nối`