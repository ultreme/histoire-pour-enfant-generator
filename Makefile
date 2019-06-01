say:
	@go run ./cmd/histoire-pour-enfant-generator/main.go | tee /tmp/a
	@cat /tmp/a | say

book.md:
	go install ./cmd/histoire-pour-enfant-generator
	echo "_Histoires pour enfants generator_" | tee book.md
	echo "_mais fait avec beaucoup d'<3 par un humain (**Manfred Touron**)._" | tee -a book.md
	echo "" | tee -a book.md
	echo "Dessin de couverture fait par tonton [Gajeb](https://gajeb.cool)." | tee -a book.md
	echo "" | tee -a book.md
	echo "_Code source disponible librement sur [moul.link/histoire-pour-enfant-generator](https://moul.link/histoire-pour-enfant-generator)_" | tee -a book.md
	echo "" | tee -a book.md
	echo "_Blog de l'auteur: **[manfred.life](https://manfred.life)**_" | tee -a book.md
	echo "" | tee -a book.md
	echo "_Editions: **Ultreme Publishing**_" | tee -a book.md
	echo "" | tee -a book.md
	echo "---" | tee -a book.md
	echo "" | tee -a book.md
	for i in {1..100}; do \
	  printf "# " | tee -a book.md; \
	  histoire-pour-enfant-generator | sed G | tee -a book.md; \
	  echo "" | tee -a book.md; \
	  echo "---" | tee -a book.md; \
	  echo "" | tee -a book.md; \
	done

book: clean book.epub

clean:
	rm -f book.epub book.md

book.docx book.epub: book.md
	pandoc book.md --epub-cover-image=cover.jpg --epub-metadata=metadata.xml -o $@
