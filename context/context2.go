func main(){
	ctx := context.Background()
	ctx = context.WithValue(ctx, "RequestID", "1145")

	processValue(ctx)
}

func processValue(ctx context.Context){
	reqID := ctx.Value("RequestID")
	fmt.Println("RequestID", reqID)
}