<script>
	import banner from '../assets/banner.png';
</script>

<svelte:head>
	<title>SplitBill | Split The Expense</title>
</svelte:head>

<main>
	<div class="lg:grid lg:grid-cols-2 lg:mx-64 md:mx-20 mx-8 lg:mt-20 mt-10">
		<div class=" my-auto text-center lg:text-left">
			<h1 class=" lg:text-5xl md:text-3xl text-2xl">Welcome to SplitBill!</h1>
			<div class="h-[1px] m-auto my-3 lg:mx-0 bg-red-400 w-10"></div>
			<h2 class=" lg:text-2xl text-sm md:text-base mt-3 leading-6 lg:leading-10">
				Share the Moment, Split the Bill <br /> Effortless Expense Management with SplitBill!
			</h2>
			<a href="/signup">
				<button class=" bg-red-500 active:bg-red-400 text-white py-2 px-3 rounded-sm mt-3">
					GET STARTED FOR FREE <span class=" animate-fade-right animate-infinite font-bold">></span>
				</button>
			</a>
		</div>
		<div><img src={banner} alt="banner" class=" lg:ml-16 lg:mt-0 mt-10" /></div>
	</div>
</main>
