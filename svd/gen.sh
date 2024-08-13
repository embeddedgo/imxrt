#!/bin/sh

set -e

cd ../../../embeddedgo/imxrt/hal
hal=$(pwd)
cd ../p
rm -rf *

svdxgen github.com/embeddedgo/imxrt/p ../svd/*.svd

for p in aipstz ccm ccm_analog dma dmamux gpio iomuxc iomuxc_gpr ocotp pmu src lpi2c lpspi lpuart usb usb_analog usbphy wdog; do
	cd $p
	xgen -g *.go
	GOOS=noos GOARCH=thumb $(emgo env GOROOT)/bin/go build -tags imxrt1060
	cd ..
done

perlscript='
s/package irq/$&\n\nimport "embedded\/rtos"/;
s/ = \d/ rtos.IRQ$&/g;
'

cd $hal/irq
rm -f *
cp ../../p/irq/* .
perl -pi -e "$perlscript" *.go
