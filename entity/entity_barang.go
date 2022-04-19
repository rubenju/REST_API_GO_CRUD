package entity

type Barang struct {
	Id    int32  `json:"p_id_barang"`
	Nama  string `json:"p_nama_barang"`
	Harga string `json:"p_harga"`
	Stok  string `json:"p_stok"`
}

type M map[string]interface{}
