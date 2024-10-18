# Fan Out

This pattern is used when multiple functions read from the same channel. The reading will stop only when the channel is closed. This characteristic is often used to distribute work amongst a group of workers to parallelize the CPU and I /O