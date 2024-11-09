-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Máy chủ: localhost:3307
-- Thời gian đã tạo: Th10 01, 2024 lúc 07:35 AM
-- Phiên bản máy phục vụ: 10.4.32-MariaDB
-- Phiên bản PHP: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `nhacungcap`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `lienlac`
--

CREATE TABLE `lienlac` (
  `id` int(11) NOT NULL,     -- Mã liên lạc 
  `name` varchar(50) NOT NULL,-- ten
  `address` varchar(50) DEFAULT NULL,-- diachi
  `gender` varchar(50) NOT NULL,-- gioitinh
  `company` varchar(50) DEFAULT NULL,-- congty 
  `state` varchar(50) NOT NULL,-- trangthai
  `email` varchar(50) DEFAULT NULL, -- Email liên lạc
  `phone` varchar(14) DEFAULT NULL,-- sdt
  `department` varchar(50)  NULL,-- phongban
  `position` varchar(50) DEFAULT NULL,-- chucvu
  `inserttime` date NOT NULL,-- thoigianthem
  `updatetime` date DEFAULT NULL-- thoigiancapnhat
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `lienlac`
--
ALTER TABLE `lienlac`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `lienlac`
--
ALTER TABLE `lienlac`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
