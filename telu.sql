-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Waktu pembuatan: 03 Jun 2022 pada 05.33
-- Versi server: 5.7.34
-- Versi PHP: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `telu`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `admins`
--

CREATE TABLE `admins` (
  `id` int(11) NOT NULL,
  `nama` varchar(250) DEFAULT NULL,
  `email` varchar(125) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `admins`
--

INSERT INTO `admins` (`id`, `nama`, `email`) VALUES
(1, 'admin', 'admin@mail.com');

-- --------------------------------------------------------

--
-- Struktur dari tabel `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) NOT NULL,
  `name` longtext,
  `address` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `customers`
--

INSERT INTO `customers` (`id`, `name`, `address`) VALUES
(1, 'asto', 'pwt');

-- --------------------------------------------------------

--
-- Struktur dari tabel `dokters`
--

CREATE TABLE `dokters` (
  `id` bigint(20) NOT NULL,
  `nama` longtext,
  `email` longtext,
  `spesialis_id` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `dokters`
--

INSERT INTO `dokters` (`id`, `nama`, `email`, `spesialis_id`) VALUES
(1, 'asto', 'asto@mail.com', 1),
(2, 'dokter faiz', 'faiz@mail.com', 1),
(3, 'dokter faiz', 'faiz@mail.com', 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `jenis_penyakits`
--

CREATE TABLE `jenis_penyakits` (
  `id` bigint(20) NOT NULL,
  `nama` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `jenis_penyakits`
--

INSERT INTO `jenis_penyakits` (`id`, `nama`) VALUES
(1, 'gigi');

-- --------------------------------------------------------

--
-- Struktur dari tabel `pasiens`
--

CREATE TABLE `pasiens` (
  `nik` varchar(191) NOT NULL,
  `nama` longtext,
  `alamat` longtext,
  `no_hp` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `pasiens`
--

INSERT INTO `pasiens` (`nik`, `nama`, `alamat`, `no_hp`) VALUES
('001', 'pasien asto', 'purworejo', '08138928201'),
('002', 'faiz', 'banjar patroman', '08881'),
('003', 'faiz', 'banjar patroman', '08881'),
('004', 'luhil', 'bundang', '082341');

-- --------------------------------------------------------

--
-- Struktur dari tabel `registrasis`
--

CREATE TABLE `registrasis` (
  `nik` longtext,
  `admin_id` bigint(20) DEFAULT NULL,
  `dokter_id` bigint(20) DEFAULT NULL,
  `jenis_penyakit_id` bigint(20) DEFAULT NULL,
  `keluhan` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data untuk tabel `registrasis`
--

INSERT INTO `registrasis` (`nik`, `admin_id`, `dokter_id`, `jenis_penyakit_id`, `keluhan`) VALUES
('001', 1, 1, 1, 'gigi senad senud');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `dokters`
--
ALTER TABLE `dokters`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `jenis_penyakits`
--
ALTER TABLE `jenis_penyakits`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `pasiens`
--
ALTER TABLE `pasiens`
  ADD PRIMARY KEY (`nik`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `admins`
--
ALTER TABLE `admins`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `dokters`
--
ALTER TABLE `dokters`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `jenis_penyakits`
--
ALTER TABLE `jenis_penyakits`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
