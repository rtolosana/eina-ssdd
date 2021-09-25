gnuplot -persist <<-EOFMarker
    set style line 1 \
        linecolor rgb '#0060ad' \
        linetype 1 linewidth 2 \
        pointtype 7 pointsize 1.5
    set style line 2 \
        linecolor rgb '#dd181f' \
        linetype 1 linewidth 2 \
        pointtype 5 pointsize 1.5
    set xlabel 'Identificador de Peticion, Linea Temporal de Ejecucion'
    set ylabel 'Tiempo de Ejecucion (segundos)'
    f(x) = 2.3
    plot "output.txt" using 1:2 title 'escenario-1' with linespoints linestyle 1, \
         f(x) title 'QoS deadline' with lines linestyle 2
EOFMarker
