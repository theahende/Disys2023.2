if [ "$*" == "" ]
then
    echo "Please enter a filename without its extension";
    exit 1;
else
    cp -r ./images ./out/;
    pandoc $1.md -o out/$1.tex --template=template/eisvogel.tex --listings -V titlepage:true -V titlepage-background:images/background9.pdf --resource-path=. --self-contained;
fi
