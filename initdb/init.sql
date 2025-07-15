CREATE TABLE dns_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL,
    value VARCHAR(255) NOT NULL,
    ttl INT DEFAULT 300
);

INSERT INTO dns_records (name, type, value, ttl) VALUES
('dominikmatic.com', 'SOA', 'ns1.dominikmatic.com. admin.dominikmatic.com. 2025071401 7200 3600 1209600 3600', 300),
('dominikmatic.com', 'NS', 'ns1.dominikmatic.com.', 300),
('dominikmatic.com', 'NS', 'ns2.dominikmatic.com.', 300),
('dominikmatic.com', 'A', '1.2.3.4', 300),
('www.dominikmatic.com', 'A', '2.3.4.5', 300),
('*.dominikmatic.com', 'A', '3.4.5.6', 300),
('dominikmatic.com', 'TXT', 'v=spf1 include:_spf.google.com ~all', 300);

