#  ____  _  ______ ____  _   _ ___ _     ____  
# |  _ \| |/ / ___| __ )| | | |_ _| |   |  _ \ 
# | |_) | ' / |  _|  _ \| | | || || |   | | | |
# |  __/| . \ |_| | |_) | |_| || || |___| |_| |
# |_|   |_|\_\____|____/ \___/|___|_____|____/ 
#                                             
# Maintainer: ohSystemmm
# Contributor: Y2kun 

pkgname=ani-track
pkgver=0.5.0
pkgrel=1
pkgdesc="Your simple cli based tool for tracking anime and manga progress."
arch=('any')
url="https://github.com/ohSystemmm/Ani-Track.git"
license=('GPL3')
depends=('make')
makedepends=('go')
source=("git+https://github.com/ohSystemmm/Ani-Track.git")
sha256sums=('SKIP')

pkgver() {
    cd "$srcdir/$pkgname"
    git describe --tags | sed 's/^v//;s/-/./g'
}

build() {
    cd "$srcdir/$pkgname"
    make install
}

package() {
    cd "$srcdir/$pkgname"
    make DESTDIR="$pkgdir" install
}