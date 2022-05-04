package entity

type Pembeli struct {
	Id      int32  `json:"p_id_pembeli"`
	Nama    string `json:"p_nama_pembeli"`
	Jk      string `json:"p_jk"`
	No_telp string `json:"p_no_telp"`
	Alamat  string `json:"p_alamat"`
	Foto    string `json:"p_foto"`
}
