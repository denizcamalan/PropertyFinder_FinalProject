Use pf_project;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;

CREATE TABLE `cart` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(250) NOT NULL,
  `vat` float UNSIGNED NOT NULL,
  `price` float UNSIGNED NOT NULL,
  `description` varchar(250) DEFAULT NULL,
  `quantity` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `products` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(200) NOT NULL,
  `price` float UNSIGNED NOT NULL,
  `vat` float UNSIGNED NOT NULL,
  `quantity` int(10) UNSIGNED NOT NULL,
  `description` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `products` (`id`, `name`, `price`, `vat`, `quantity`, `description`) VALUES
(1, 'Apple 13 Pro', 1025.6, 0.08, 23, 'phone'),
(2, 'Arcelik Hg567', 235.3, 0.08, 12, 'refrigerator'),
(3, 'Beko Ecological Inverter', 510.5, 0.1, 18, 'air conditioner'),
(4, 'Hyundai Inverter', 1400.9, 0.08, 5, 'generator'),
(5, 'JBL', 80.7, 0.1, 39, 'headphone'),
(6, 'Macbook Air', 1230.6, 0.18, 20, 'computer'),
(7, 'Samsung S5', 650.4, 0.18, 14, 'phone'),
(8, 'Vestel Smart', 350.2, 0.1, 30, 'tv');
