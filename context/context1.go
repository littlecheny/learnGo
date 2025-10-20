func fetchData(ctx context.Context) (string error){
	ctx, cancel := context.WithTimeout(ctx, 2*Time.Second)
	defer cancel()

	ch := make(chan string,1)
	go func(){
		Time.Sleep(2*Time.Second)
		ch <- "data"
	}()

	select{
	case <-ctx.Done():
		return "",ctx.err()
	case result <- ch:
		return result,nil
	}
}