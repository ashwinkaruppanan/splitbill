<script>
	import logo from '../assets/Split.png';
</script>

<header class=" h-20 w-full shadow-md flex justify-between items-center">
	<div class=" lg:ml-64 md:ml-20 ml-8">
		<a href="/" class=" flex items-center">
			<img src={logo} alt="logo" class=" sm:h-12 h-8" />
			<h1 class=" ml-2 sm:text-2xl font-semibold">
				Split<span class=" text-red-500">Bill</span>
			</h1>
		</a>
	</div>
	<div class=" lg:mr-64 md:mr-20 mr-8">
		<a href="/signin">
			<button class="mr-4 hover:text-red-500"> Sign in </button>
		</a>
		<a href="/signup">
			<button
				class="  px-2 py-1 ring-1 ring-red-500 rounded text-red-500 hover:bg-red-500 hover:text-white active:bg-red-400"
			>
				Sign up
			</button>
		</a>
	</div>
</header>
