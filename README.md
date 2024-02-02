On my journey of figuring out concurrency in golang, I thought of implementing the Therac-25 incident, an gruesome reminder of what can go wrong when software engineering goes awry. This code uses goroutines to implement the concurrent nature of 
setting radiation modes, waiting 8 seconds before setting mode completely and a blocking call to irradiate the patient for 10 seconds. 
