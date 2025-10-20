func (w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	resultchan := make(chan string)

	go func(){
		sleep(2*time.Second)
		resultchan <- "result"
	}

	select{
	case <-ctx.done():
		http.Error(w,"request cancled",http.StatusRequestTimeout)
	case result := <- resultchan:
		fmt.Fprintln(w,result)
	}
}
