package entity

type Transaksi struct {
	Id_transaksi int32  `json:"p_id_transaksi"`
	Id_barang    string `json:"p_id_barang"`
	Id_pembeli   string `json:"p_id_pembeli"`
	Tanggal      string `json:"p_tanggal"`
	Keterangan   string `json:"p_keterangan"`
}
