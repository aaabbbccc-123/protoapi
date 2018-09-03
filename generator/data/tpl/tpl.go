// Code generated by "esc -o generator/data/tpl/tpl.go -modtime 0 -pkg=tpl generator/template"; DO NOT EDIT.

package tpl

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/generator/template/echo_enum.gogo": {
		local:   "generator/template/echo_enum.gogo",
		size:    352,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/3yOsWrDQBBEa+1XDCaFBLHcO7iKE0hjB2LSGBcbaSNErNVxOhXm2H8PkmKBmnSPuZnb
t9nguS0Flah4DlLi6wbn29Cyq5+wP+JwPOFl/3bKiRwXP1wJYszfJzQjCjc3RgduxAy1BqKi1S4gJQCI
cQ3PWgny11quZQezv4f7JsaHO+6G+JOvvcytNURLM8qIvnstkBaD8bzO8BF8rVWaoRsBkRLlRjpsd2jY
nefqZSrE8eN/3RZ+W6xmXj0uxpNZYkSJl9B7xXj5PCheyOg3AAD///CkU9tgAQAA
`,
	},

	"/generator/template/echo_service.gogo": {
		local:   "generator/template/echo_service.gogo",
		size:    1543,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xUT0/bThA9ez/F/KyfKhsFh7a3IC5ASmkFQSR32NiTZIuz6+6OaZG1372atWMSQnqo
T9bMmz9v5u0Mh3BhCoQlarSSsID5C1TWkJGVOoXLCdxOZjC+vJ5lQlQyf5JLhKbJ7tpf74VQ68pYgkRE
sUYaroiqWIgoXipa1fMsN+thKeeOZP40xHxlYpEKMRxyllu5Ru9BOaAVgtKEdiFzhNxokko7kGUZXGyw
pizROkEvFW4H91GNiJrmGKzUS4TsBmllCgfeCwDggJmiEr1Pjpomu9ZVTbOXCr1PIVgmNfWmplEL0AjZ
2Fpj2QZx7P0AGNjbGIe68D7tKhwD6oILetF2GaDX3KCWJTiydU7cZhg5APf++RO03+MPZ/Qozk2B8aOI
btA5nrUjq/RyB7JuXYy6RJKqdK9DaPwGVbSu+LHv5h5dZbTDdxrauHYSvVa0nZtLBk7sO9pl10GRjW3R
95exqHUOD/3+Hr5KXZRoE2efoWn+78wpsFSyzvmFYxoRWaTaauAUSd4iLowm/E0pJGgthOopQyOlYXQG
Gn8lb9Yt2LlgKJxBnp0rXSRKp6fB8t8ZaFWGBBGz5hx7cwvedhAj+LAzhoZXO4KPJycnA+h2OOLMrWyS
1A842Lf5A5s8+zad3Cb8brIpSarduSzu8WeNjgbATaSCI0QUmZoOSjNQ7wTJXTv7nG2JXulA/BCnQGlj
HIGpKfTJK3ynHu/xdQIAb4fQyXLUNuX7VKG5DZmD9Cffe9r8lLbeFZ+Ne1wqR2h3zkftsAAyMFe6AGtq
4kMRlLYHTxCOgnLG+coMoJVdr7qmf8qHzkg/kCnaZ5VjO5K7yXQWb0CYXY1nSRzOJK28jwcHBJ/uXY6/
Frkab9fgkv9WZP9a/QkAAP//uFGvAgcGAAA=
`,
	},

	"/generator/template/echo_struct.gogo": {
		local:   "generator/template/echo_struct.gogo",
		size:    235,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/1zNsa7CIBTG8Z2n+NL9wn7v2F4Tl9ahD1Asx6baAgIdCOHdDakx6vbPj/AdIVAbRZhI
k5OBFM4R1plgpJ3/0HRoux7/zbHnjFk53uRESImf9syZMSEK1Iv0vpVroRAtfRl8cNsYkBgApPQDJ/VE
4IeZFuWR8/OB93NYyoeS0ZYart7o3yolvm9VuBi3fsB9IxffZXjdIa3KemaPAAAA//9AUK296wAAAA==
`,
	},

	"/generator/template/php.gophp": {
		local:   "generator/template/php.gophp",
		size:    2425,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5xVTW/jNhC961cMBAOWAlvrbQ9dxI2LNEm3OSQput6TEwi0NLG5kEmFHCbrCPrvBfVl
SZa3RXkwTHjmzeObmedff0u3qfPh7MyBM1huuQaugcEzTxA2KFAxwhjWe0iVJMlSDt4rKs2lCIx5N0Ek
dx/qn3yL8blJYnQOWQbBku8Q8tz+eP0A9w9LuLm+XQYOnH1wHMF2qFMWIWRZcM92+MVe8nzuOEYj3HyP
MCUuxby4fjbv7wn+SZQ+XiUcBQHTYK/lbe5kmWJigxDcodZsgzrPnShhWkOWEacEoSiS507mAABk2RSq
jD84JrEuiAJAqvgrI4RRRcsyKhNQxHWMWSc8gmcjIksRXlnCY0boMaXYHkYKdSqFRr/MLD5/WNQe/gwe
1xrJa/JXbkPCffL9FlKJZnO4flh/w4gguGbElvvUSj6iLdfTRZMNFyDwrS/FqULzLANMNA7iDOcUKY0+
9uQFRI8ybZV8K6g0/fXccYMztjMoJAF+55pcf35A62h4qJRXvRluZa9NGilsanlZNqhenvdUqjt/GIh+
W4dU6k1P7gwy2nQY9XEVklHiCH7eefnpqSQZFvN4Cnf13+ayDDihVWsG4GJxRHW6OJCY9BDLCfsXgEn7
gcftt+ep1iN36uB68bsrb90KI2tQo23LOQa1C8NICk3KROSN1kzjV8XhAsYff/olmAWz4OP5p9mn2fjE
JBzgq8U7OJXXUWHVudkztsVCo/i4kKMqPTmOI75DaagI+3nWDXhqbn53/Cpva54ZsSS5TLk1gpcJjHZI
WxlPYGQUn8BIGkoNWQ2P3hkzYnDR4z/+pqUoeSt8Ge7902GnGyex+9IXbrpQ+GJQk9dnZUv7bRT2diUF
lWI3mNPFBul3Ge89v/hahWjPrzpe5EZNoqUexhjJGL0W5ARIGWyVa7l0lbxyUSmpBgx6yO2OslburvzH
ssY74HdDBetHDtS0ApiknrtWC70Weitwumj+vAbxu0m1I5W58/9h9ldMWIN/5iKGpv9rGe/PwYWg3c0j
NXpef1eMhT5hfm2fD25Faqg0rGI0f2y03aWwBmXX947iPHfL+9e/b4vvYXh/eXfz5a/Lq5swhADcx8cs
Cx4KzSt/bBaw8qbc+ScAAP//SLLRKXkJAAA=
`,
	},

	"/generator/template/spring_service.gojava": {
		local:   "generator/template/spring_service.gojava",
		size:    854,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/6yRv27CQAzG9zyFlSkMPR4gC6IgxABBVdTdSUw4Ab7rnSOEonv3KuFPGNpKRWSK7vNn
fz97PIZ3UxHUxORQqILiDNYZMWh1CrMM1lkO89kyV1FksdxjTdC2anP5DSGNIn20xgkYVytvneZ66/BI
J+P26kSFKjRXCpmNoGjDakGyQms11+l/rRvjn/Z+kLeGPU1NdX7C/NWQl4s3sk1x0CVg4cVhKVAe0Ptu
K2s8UghT9ARtBADQtm/gkGsCtSLZmcpDCHdFb4EJVGb7GZ94aAjixTyPb0WTB+Ak7tcuuxDi0UV9ROpf
rsHaVmWN2Ebys6UQhmRdu2TyANNJSx4qNY+uybvPkTSOB3uieZT26sBAXHVp/2YagIbjv4JnQZK8lOB+
0l9H/jAvvba6dQrRdwAAAP//oNlDdVYDAAA=
`,
	},

	"/generator/template/spring_struct.gojava": {
		local:   "generator/template/spring_struct.gojava",
		size:    585,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5SPwWoCMRCG73mKOdqD8QGWQkFbqBT1sC8w7o7baDYJk1mphLx7ibqtUmhpTjOB+f7v
n81g7luCjhwxCrWwPUFgLx6DqWCxhtW6hufFa62VCtgcsCNISW8uY86VUqYPngUa3+sdRiH+6K3eY3OI
3ml0zguK8U4vo3dzJhTP1b+ONuwDsZy+s/Z4RD2IsfrNRKmUCsPWmgYaizEWv3kZVthTzpAUAEBKU2B0
HYF+MWTbmPP5P7A5ohDsjENbTpd4xPoUzpcJdIFA6TlCyLWQszrvTzedLrSLx73BpKzeCQ+NbJCxz/nh
D6vy5N1EfaMAj78KlfWqldIdE6Zj1y+7m5IdScHWRmzhTka18phkYPczNl9jxuysPgMAAP//ioLHmkkC
AAA=
`,
	},

	"/generator/template/ts/data.gots": {
		local:   "generator/template/ts/data.gots",
		size:    606,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5yPQWviQBiG74H8hw8vQlj1vrAHF7OwIIssspdlWcbkUwfiJDszWVaGgYVaWqEWodqD
h/ZUEAT1VGjtzzGJP6PEkIqHXswpw/u8z7xTsSzTsKDZpQLa1EOgAjrIkBOJLrT6UAy4L30S0OIx5/hM
EsoEEM8D2UVwiSQgJA8dGXKEFlLWgVCgC5TtgYNWChDI/1IHRSotnfKlxd3qIb692G4ek8l9fDnePl/n
Y9M0i6Kr82i8qDa+RsNRsljt1mfJZB4P/0cv02QyT2aDeLqOR8tkcxPfDaLlbPs0NA2rYhpKlYAT1kEo
2yzsCdDaNPBf4HMJyMIeKFX+RnqoNSjTAABQKi98oei5WQOyJEc/pf8/iBei1h/easjcPayza/Pj0Yga
kaTZD/BoCGUSeZs4eNKajykkRWo9+LVWirYB/0C5TlroQaFe/WzXf3+3G3a1adcKWv/8pRQyNxe+/4DX
AAAA///KSGFSXgIAAA==
`,
	},

	"/generator/template/ts/helper.gots": {
		local:   "generator/template/ts/helper.gots",
		size:    2447,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5RWbW/bNhD+bsD/4SZslVS7UrZ1QKFMTV/WbR2GuouTT4Y/MNLJZk2TGkUp0dr894Gk
JEuxhaYBAoh3zz33yqPDp0+nEzD/8ConkuyhIgwIFEpSvrFyiaqUvADCAXkiUkwbNagtUZAQDjcIJM+R
p6AEkJxCKZk2DqcTvMuFVJCVPFFUtBReRVjU0PjtB3yeTgAArMMGeX35/q3Y54IjV9rKDyzIAnNGEvTC
H56fhRs6B/eVe1r/82urj0b0Pz0PN3Nwvx9Tv7Xm8zH9mTGfjah/eWPNV2P636x+7frn08n9dDKdNJ2B
NyVlKRC4vvwbbuqmyrpWpluFrrfaIiBP+00sJYPIKG5IgeaYCWkEBcqKJthHN1TWQOK/JRYKxM0nTFQA
GGwCY/wnMiYuG+2tKFmq265NLBREBqrOcQAcTFAE3ShBVjIG15K1rX9mfPzx7kpnucM6rAgrEXJCZWFJ
8I7sc4aROem0tHUMzlapPApDJhLCtqJQ0YuzF2eOQRG5gRg+c7LHCJxb5JtbpM4cOE12VqAo4c69Abex
xS/hJOWFpok7lieaJO44RqZ9gxwlUfhPibK+luzXq5deKbvRnze1j+Dq6BbQDLzvrNpvZb3bUUp2boVm
XvQHQ6X5VAExrNbnWmwVurSe1u6wBsrhBKvWVqSlbP0v7BDkUiihextsSbG45R+lyFGq2tth7Q9o9J/e
IHHjY7XDet0j7WJtPRhwHAPX8/Dli5kgkUErdkueYkY5pu6RH1uIMXKT7lwTFT1IGGqvtAAiJamPQgmU
WJomeL51v2pm+7WGr49j2EFsijoDd7V2+7EAsgJH4Ef1Mh2rCBstlcYEmZDvSLL1Kj2jD5i7xFKicKgy
uY1m9jtleCIx4xRi0Ibvl4vW9vwY1Xm2hENAUwYdQttZ696Cx/3+tVx8COyFoFntVQ893w+PZuyDvCy2
XvPA7HzdlNiFWffk+H6vuv7h9rTjUqCkhNH/MP1oN2Lc8H4SlHvuE7OdD5fzIXyQi165sxj0ZQ8oT/Fu
kXnuhWuL/+xHuAD3woUINCvMjlwPgztc+eHr8HDlP3bhN6IPZI99COgN18e1a6wDdnvtgGzW8jds5b73
GBzzXCytzAIGbmNwlqQ2IOdRezrsEwat7Vf286nV3IvzIOzH9rUfMGYGwAkd2+Au6Rk4gTMbUNln//8A
AAD//8c9h/ePCQAA
`,
	},

	"/generator/template/ts/ts_service.gots": {
		local:   "generator/template/ts/ts_service.gots",
		size:    2675,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/8RUbW8bRRD+bsn/YWRV8oucuwaJqrgEKYKC8gEaqeEHbO7G9pbz7nV3L210nNRS0lIp
bSI1qoAEWgRFFU1ikKAkaUP/zJ3tfOpfQHvr80uTqKkqxH2wfDvPPDPz7DNnVyr5XAXmmlRCnXoIVEID
GQqi0IX5RSj6gitOfFo0ODQwhzNFKJNQF5wpZC5Mz86Aw10E1SQKrnDxBVyhqgmqieDReUHEIhTJVcpl
saoPBda5wCpQVZQg8HJABbomuQ/TvVAmFfE8dIGylMoX/BI6Kmtm2Gtam0q4IqhSyDR+btHHi46g/gCe
gnzBF6iLEgjME0kdCCRpINS5MEMQzwPCXGiRRWCILhD3UiBVC5kC4jhcuJQ1QHGQPjq0Tp2sqWwMg2Qu
SKoCoihnuv7EyR8N77Ufde7fip897a496HyzGu/dzW5CR00oWV5KVp8kt+90n7R7v9/orj2enp3pfv91
/Ozn7sPrL58vJ7tP4/0X3bXH3c3NeOd25/5usncvVffl82VI1h92tn452LjW+/V6/OKHXvt6Gkq2vk02
Hsd7dw9+2u2ub8c7WyMVby4ZckM7rnCv/cg028eu/NZZWY3/WTcNTs/OmB6TB3udjc1hj9s/Hny31Lmx
lNz8K1lp927sm446D3c7d7aTpb/j/XumE/Nfh/74Kt65k6wu964tx/sbya29zvqfnbXdfK5i53O05XOh
IJ2lCiFM6z+zgreoRIi0Y1uZxc4N0GE+BwAQhoKwBsIptehjFU7Nc+5BbQpKDVQzKfIjoogeW4L1ccAc
fbuyHEVZ+oRJhSiqQnaGzNWArLRlu0SR0doDG38uPBjCmuj5KDQwn7Mr80Sm8WSlbabVszqcSQVZaAoK
TaX8mm1PvveONXnmrDU5+a515nTt7Omzpwvn8jnbhkCikUazpv06HpHyM9JCPaj14eBtQjcdhnbF4Iww
w6FNdAJOtVA1uauTCz6XqpAlTgCtA15OtbuIYoE6+KlywbrgG9Gg8Mn5uRH0KFEDhzx9+VIICsGFRqSs
5/WbvowRUg3Fq6mu9X6rEIaWniiKSj4RpCVr+mSG+YHSyVFUrkHfH++HoXUhUIMIfAkMF1B8kBnEQwWB
8GogldCfganRuyv1L6IKhTAc6hpFhfSk30ShfM5w2TYolMq8CFSBYOZqrDDsixFFpUDzhSGtA8OBRkag
KDLzhCF6EqMoBPMOUV+1suHWj6WayEoCJUwNhske24YmYa6HkOo7HqR10GmWNq2Vxsuv5h8Ns1xUhHry
SPjIyH3pLYH6I3ocyWGK6PDRSRiBSAhD46QxhY7gHAojA8dBKUHzjGMO1ZTcW8BhUYHS50xvna475q6R
4pE2xajdK7b+7S9sHZXTfM0a/g+rMZTxv94TF+sk8BTMXrg4N7YvqTKlENItMctRg6L+DhWr6WXVBjtR
Pm4FjrX/Caz/hrZ/C8tHr7Hdm1g9Omr4oy3+9vZ+xdr/BgAA///CGezJcwoAAA==
`,
	},

	"/generator/template/ts/ts_service.govue": {
		local:   "generator/template/ts/ts_service.govue",
		size:    2545,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/4xWX28Txxd9t+TvcGUh2Y6cdV5+EjI/UCnQKlLbRCXwPtm9tgfWM8vM3YRoOxKUBooU
IBIRapu0ULVUqPxJK7WUBFK+jNc2T3yFajz+s04TyjzZd86998w5d8auTk3lc1Ow0OQa6jxE4BoaKFAx
wgAWV6AYKUmSRbzocOhgvhTEuNBQV1IQigBOzs+CLwMEajKCZakuwjKnJlATIeSLiqkVKJ6P8XPUMlY+
Fit2S2FdKqwAp6IGhZdirjBwJSbAlhcXmlgYYgBc9MtGSl5An4bExrz7PLiGZcWJUFj8wkqEZ33FoxG8
D4qUXOIBamCwyDT3IdasgVCXyh2IhSEwEUCLrYBADIAFF2JNLRQEzPelCrhoAEnQEfq8zv0hqeFhHFIE
oDnFjLgUtv/0+y8L720/7Ny70X75vLtxv/P1env39tAVu+u20rXVdP1xevNW9/F277dr3Y1HJ+dnu999
1X75U/fB1bev1tKd5+29192NR90nT9ovbnbu7aS7dzMav321Bunmg87Tn99sXen9crX9+vve9tUMIH36
Tbr1qL17+82PO93NZ+0XTzPdr6+6Rq7FpNq97YeO+AB759fOnfX235uO7Mn5Wcc3vb/b2Xoy5vvshzff
rnaurabX/0zvbPeu7TlenQc7nVvP0tW/2nt3HRP32W79/mX7xa10fa13Za29t5Xe2O1s/tHZ2Mnnpqr5
HG9FUhEkcEq2IilQUAXOxwjGDnELiksxTkdKRqhoZTpAXypGUhWPjTIzWmRS1HBGx8AknwMASBLFRAPh
CK1EWIEji1KGUDsOpQbSbB95mhGzWmnwPoqFb8dDl40Zpk+7VDCmAsMYisAChqS9asCIZXuP7sE5FcIY
1sQwQlXM5/K58zF6scZS5jzlY3bjg5Ey+Rxe7pcLsM7ikMAPmdaQJN4p++Ez1kJjAC/bq6/7Kg7OXK3a
t0GTin2SysUygVJ5CLRLxxGqUtkFjGUwKPEh0wjnVOgCkeJLjNDeUHuoGmhS9todh0KTKKpVq6H0WdiU
mmpHZ47OFDL6+UO6VvgxeZjOquxsGlsw4UALqSkDm16IpKbCZCqvA17qO3oW1RL38VMKwJuLnJVQ+PjM
wr6MbMEGTtYbmDuColJSWWS/wxn7zY5LpsEY7jlTShFTrKVrNjIrophsgjHlGswr2eIaP+EX8f9J4s3F
NNqFL6Bf/IR1Z+yPXSESxBOiZ+arRE2uvYEvlUKSjPU2plCBwohWwU1YtnK1Cgtzp+dqECChanGBEGv7
jFPVCj35vtrnFDVNVlBIsbK/BVx7R+woeEkykNeYUqzCysgkgSPhnerGOKGSBEONxiTuqxmYUPaoiaI0
2c/11HD8RHaKs8uOugzRC2WjpFDbQx8Eq1bBWdtkIgi5aBwM43WwVTx7xb1+QvmwxgfDvQCJ8VC/My2j
5GBEPIX2N+ywYoeXModvvU8HYPaJcVNvzCF9DukxqD8qqFBHUmh0NSem/QBTTOXfMUfJem39kvbfiotp
0LHftIX/NzNTASFBINl/PP89E87EY+88gsBldxtLhbPSzqy9AstKikbhoMx9emQh+98Vk8/9EwAA//+d
KTPT8QkAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/generator": {
		isDir: true,
		local: "generator",
	},

	"/generator/template": {
		isDir: true,
		local: "generator/template",
	},

	"/generator/template/ts": {
		isDir: true,
		local: "generator/template/ts",
	},
}
