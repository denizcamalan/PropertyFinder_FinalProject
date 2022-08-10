use pf_project;
SET NAMES UTF8;

-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Aug 09, 2022 at 01:11 AM
-- Server version: 5.7.34
-- PHP Version: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pf_project`
--

-- --------------------------------------------------------

--
-- Table structure for table `cart`
--

CREATE TABLE `cart` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(250) NOT NULL,
  `vat` float UNSIGNED NOT NULL,
  `price` float UNSIGNED NOT NULL,
  `description` varchar(250) DEFAULT NULL,
  `quantity` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `cart`
--

INSERT INTO `cart` (`id`, `name`, `vat`, `price`, `description`, `quantity`) VALUES

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(200) NOT NULL,
  `price` float UNSIGNED NOT NULL,
  `vat` float UNSIGNED NOT NULL,
  `quantity` int(10) UNSIGNED NOT NULL,
  `description` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `price`, `vat`, `quantity`, `description`) VALUES
(3, 'Apple 13 Pro', 1025.6, 0.08, 23, 'phone'),
(4, 'Arcelik Hg567', 235.3, 0.08, 12, 'refrigerator'),
(8, 'Beko Ecological Inverter', 510.5, 0.1, 18, 'air conditioner'),
(7, 'Hyundai Inverter', 1400.9, 0.08, 5, 'generator'),
(6, 'JBL', 80.7, 0.1, 39, 'headphone'),
(1, 'Macbook Air', 1230.6, 0.18, 20, 'computer'),
(2, 'Samsung S5', 650.4, 0.18, 14, 'phone'),
(5, 'Vestel Smart', 350.2, 0.1, 30, 'tv');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cart`
--
ALTER TABLE `cart`
  ADD UNIQUE KEY `id` (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD UNIQUE KEY `name` (`name`),
  ADD KEY `id` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
