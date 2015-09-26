echo "Removing stale dependencies ..."
rm -rf pkg
echo "Installing dependencies ..."
mkdir pkg
cd src/utils/parallel
go install
echo "\t... Installed parallel √"
cd ../../hashing/lsh
go install
echo "\t... Installed lsh √"
cd ../../projection/rp
go install
echo "\t... Installed rp √"
cd ../fjlt
go install
echo "\t... Installed fjlt √"
cd ../../counting/counts
go install
echo "\t... Installed counts √"
cd ../../probing/gb
go install
echo "\t... Installed gb √"
cd ../../
