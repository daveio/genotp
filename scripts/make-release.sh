#!/bin/sh

BINARYNAME="gotp"
PACKAGENAME="github.com/daveio/gotp"
SCRIPTDIR="$(dirname $0)"
SRCDIR="${SCRIPTDIR}/.."
RELEASEDIR="${SRCDIR}/release"

if [[ ! -d ${RELEASEDIR} ]]; then
    mkdir -p ${RELEASEDIR}
fi

platforms=(
    "darwin/amd64"
    "linux/386"
    "linux/amd64"
    "windows/386"
    "windows/amd64"
)

for platform in "${platforms[@]}"
do
    echo $platform
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    if [[ ${GOARCH} == "386" ]]; then
        RELARCH="i${GOARCH}"
    else
        RELARCH=${GOARCH}
    fi
    OUTDIR="${RELEASEDIR}/${BINARYNAME}-${GOOS}_${RELARCH}"
    mkdir -p "${OUTDIR}"
    output_name="${OUTDIR}/${BINARYNAME}"
    if [ $GOOS = "windows" ]; then
        output_name+=".exe"
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o ${output_name} ${PACKAGENAME}
done

cd ${RELEASEDIR}
for i in *-windows_*; do
    7z a -tzip -mx=9 ${i}.zip ${i}
    rm -rf ${i}
done
for i in *-darwin_* *-linux_*; do
    tar -cf ${i}.tar ${i}
    gzip -9 ${i}.tar
    rm -rf ${i}
done
