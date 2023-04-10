package model

func partition(u []Product, low, high int, typ int) int {
	pivot := u[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if typ == 1 {
			if u[j].Price > pivot.Price {
				i++
				u[i], u[j] = u[j], u[i]
			}
		} else if typ == 0 {
			if u[j].TotalRating > pivot.TotalRating {
				i++
				u[i], u[j] = u[j], u[i]
			}
		}
	}
	u[i+1], u[high] = u[high], u[i+1]
	return i + 1
}

func qSort(u []Product, low, high int, typ int) {
	if low < high {
		pivot := partition(u, low, high, typ)
		qSort(u, low, pivot-1, typ)
		qSort(u, pivot+1, high, typ)
	}
}
