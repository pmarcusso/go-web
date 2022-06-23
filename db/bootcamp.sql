USE bootcamp;
CREATE TABLE `transactions` (
                                id      INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
                                cod_transaction	INT(10) NOT NULL,
                                currency_type    VARCHAR(100) NOT NULL,
                                issuer   VARCHAR(100) NOT NULL,
                                receiver   VARCHAR(100) NOT NULL,
                                date_transaction TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
                                PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;