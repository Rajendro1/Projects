-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Oct 30, 2022 at 05:31 AM
-- Server version: 8.0.30-0ubuntu0.20.04.2
-- PHP Version: 7.4.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `THIRDESSENTIAL`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `name` varchar(40) NOT NULL,
  `price` float NOT NULL,
  `image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `user_id`, `name`, `price`, `image`) VALUES
(3, 1, 'hairOIL', 200, '1+383589c6-c269-4921-b031-b8644729767b.jpg'),
(5, 2, 'Router', 1000, '2+5c310a33-4250-4207-a6ce-16236819b6ff.jpg'),
(6, 2, 'RouterNEW', 1000, '2+dae91630-c54f-45fc-a07c-9a70025b480c.jpg'),
(7, 2, 'RouterNEW1', 1000, '2+043795dd-d854-4283-ba4d-02b28441c4c1.jpg'),
(8, 2, 'RouterNEW11', 1000, '2+15f83c10-f511-4ebe-85f2-aef067bd38a2.jpg'),
(9, 2, 'RouterNEW11', 1000, '2+65557881-c89f-43ea-8181-c4e7d7e7d8f5.jpg');

--
-- Triggers `products`
--
DELIMITER $$
CREATE TRIGGER `deleteProducts_audit` AFTER DELETE ON `products` FOR EACH ROW INSERT INTO products_audit(`product_id`, `user_id`, `name`, `price`, `image`, `action`) VALUES(OLD.id, OLD.user_id, OLD.name, OLD.price, OLD.image,'delete')
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `insertIntoProducts_audit` AFTER INSERT ON `products` FOR EACH ROW INSERT INTO products_audit(`product_id`, `user_id`, `name`, `price`, `image`, `action`) VALUES(NEW.id, NEW.user_id, NEW.name, NEW.price, NEW.image,'create')
$$
DELIMITER ;
DELIMITER $$
CREATE TRIGGER `updateProducts_audit` AFTER UPDATE ON `products` FOR EACH ROW INSERT INTO products_audit(`product_id`, `user_id`, `name`, `price`, `image`, `action`) VALUES(NEW.id, NEW.user_id, NEW.name, NEW.price, NEW.image,'update')
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `products_audit`
--

CREATE TABLE `products_audit` (
  `product_id` int NOT NULL,
  `user_id` int NOT NULL,
  `name` varchar(40) NOT NULL,
  `price` float NOT NULL,
  `image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `action` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `products_audit`
--

INSERT INTO `products_audit` (`product_id`, `user_id`, `name`, `price`, `image`, `time`, `action`) VALUES
(1, 1, 'Laptop', 80000, '1+1.jpg', '2022-10-29 12:32:01', 'create'),
(1, 1, 'PC', 80000, '1+1.jpg', '2022-10-29 12:39:40', 'update'),
(1, 1, 'Mobile', 80000, '1+1.jpg', '2022-10-29 12:39:53', 'update'),
(1, 1, 'Mobile', 80000, '1+1.jpg', '2022-10-29 12:40:08', 'delete'),
(3, 1, 'hairOIL', 200, '1+383589c6-c269-4921-b031-b8644729767b.jpg', '2022-10-30 03:27:37', 'create'),
(5, 2, 'Router', 1000, '2+5c310a33-4250-4207-a6ce-16236819b6ff.jpg', '2022-10-30 04:48:29', 'create'),
(6, 2, 'RouterNEW', 1000, '2+dae91630-c54f-45fc-a07c-9a70025b480c.jpg', '2022-10-30 04:53:17', 'create'),
(7, 2, 'RouterNEW1', 1000, '2+043795dd-d854-4283-ba4d-02b28441c4c1.jpg', '2022-10-30 04:54:59', 'create'),
(8, 2, 'RouterNEW11', 1000, '2+15f83c10-f511-4ebe-85f2-aef067bd38a2.jpg', '2022-10-30 04:57:54', 'create'),
(9, 2, 'RouterNEW11', 1000, '2+65557881-c89f-43ea-8181-c4e7d7e7d8f5.jpg', '2022-10-30 05:00:03', 'create'),
(10, 2, 'RouterNEW111', 1000, '2+05b21cee-6c98-4270-9ffb-6e1dc36184bf.jpg', '2022-10-30 05:01:40', 'create'),
(11, 2, 'RouterNEW1111', 1000, '2+5d02f145-2de4-4529-b518-b03f4ae87122.jpg', '2022-10-30 05:04:25', 'create'),
(12, 2, 'RouterNEW11111', 1000, '2+04f04c12-e3ce-4a4e-b44d-f0f6e9f57e68.jpg', '2022-10-30 05:07:09', 'create'),
(12, 2, 'TP-link', 500, '2+ad670398-7cc3-4a79-a405-7c4e59396761.jpg', '2022-10-30 05:11:50', 'update'),
(12, 2, 'TP-link', 500, '2+ad670398-7cc3-4a79-a405-7c4e59396761.jpg', '2022-10-30 05:18:15', 'delete'),
(11, 2, 'RouterNEW1111', 1000, '2+5d02f145-2de4-4529-b518-b03f4ae87122.jpg', '2022-10-30 05:25:28', 'delete'),
(10, 2, 'RouterNEW111', 1000, '2+05b21cee-6c98-4270-9ffb-6e1dc36184bf.jpg', '2022-10-30 05:29:47', 'delete');

-- --------------------------------------------------------

--
-- Table structure for table `superadmins`
--

CREATE TABLE `superadmins` (
  `id` int NOT NULL,
  `name` varchar(40) NOT NULL,
  `email` varchar(40) NOT NULL,
  `phone` varchar(15) NOT NULL,
  `address` varchar(100) NOT NULL,
  `password` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `superadmins`
--

INSERT INTO `superadmins` (`id`, `name`, `email`, `phone`, `address`, `password`) VALUES
(1, 'Raj', 'rajandroprosad1@gmail.com', '8250771252', 'Rangilabad,West Bengal', ''),
(2, 'Rajendro Prasad Sau', 'saurajandroprosad@gmail.com', '987654321', 'Kolkata', '$2a$14$xvTG/AXYzP.TfxJf2T2MUO/.t0Gu9YAhdUTTF95xOzUDKdprLhAGG');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar(40) NOT NULL,
  `email` varchar(40) NOT NULL,
  `phone` varchar(15) NOT NULL,
  `address` varchar(100) NOT NULL,
  `password` varchar(500) NOT NULL,
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `logout_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `phone`, `address`, `password`, `login_time`, `logout_time`) VALUES
(1, 'Rjendro', 'rajandroprosad@gmail.com', '8250771252', 'vill+po-Rangilabad', '', '2022-10-29 12:30:38', '2022-10-29 12:30:38'),
(2, 'Sukannya', 'sausukannya@gmail.com', '123456789', 'West Bengal', '$2a$14$PlRCigQVF0mGZL7BpkgD9.0C7OG.7CwmHqRZEuZYq.laj6ows8CC6', '2022-10-30 04:21:32', '2022-10-30 04:21:19');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `superadmins`
--
ALTER TABLE `superadmins`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `superadmins`
--
ALTER TABLE `superadmins`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
