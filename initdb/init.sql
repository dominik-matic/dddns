CREATE TABLE dns_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL,
    value VARCHAR(255) NOT NULL,
    ttl INT DEFAULT 300
);

INSERT INTO dns_records (name, type, value, ttl) VALUES
('dominikmatic.com', 'SOA', 'ns1.dominikmatic.com. admin.dominikmatic.com. 2025071701 3600 1800 1209600 3600', 300),
('dominikmatic.com', 'NS', 'ns1.dominikmatic.com.', 300),
('dominikmatic.com', 'NS', 'ns2.dominikmatic.com.', 300),
('ns1.dominikmatic.com', 'A', '3.73.187.39', 300),
('ns2.dominikmatic.com', 'A', '3.73.187.39', 300),
('dominikmatic.com', 'A', '3.73.187.39', 300),
('www.dominikmatic.com', 'A', '3.73.187.39', 300),
('*.dominikmatic.com', 'A', '3.73.187.39', 300);
