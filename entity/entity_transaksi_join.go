package entity

type TransaksiJoin struct {
	Id_transaksi int32  `json:"p_id_transaksi"`
	Nama_barang  string `json:"p_nama_barang"`
	Harga        string `json:"p_harga"`
	Nama_pembeli string `json:"p_nama_pembeli"`
	Tanggal      string `json:"p_tanggal" form:"tanggal"`
	Keterangan   string `json:"p_keterangan"`
}
