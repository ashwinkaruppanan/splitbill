<script>
	let password = false;
</script>

<div class="flex justify-center md:mt-32 mt-12">
	<div class=" md:w-[28rem] lg:h-[25rem] mx-5">
		<h1 class=" sm:text-2xl text-xl font-semibold text-center">
			Split<span class=" text-red-500">Bill</span>
		</h1>
		<h1 class="lg:text-xl mt-2 text-center">Start with your free account today</h1>
		<form class=" mt-4">
			<label for="name">Full Name</label>
			<input
				type="text"
				name="name"
				id=""
				class="outline-none p-2 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
			/>
			<div class=" mt-4">
				<label for="email"> Email address </label><br />
				<input
					type="email"
					name="email"
					class="outline-none p-2 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
				/>
			</div>

			<div class=" relative mt-4">
				<label for="password">Password</label>
				<input
					type={password ? 'text' : 'password'}
					name="password"
					class="outline-none p-2 pr-16 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
				/>

				<button
					class=" absolute right-2 bottom-2"
					on:click={() => {
						password = !password;
					}}
				>
					{password ? 'hide' : 'show'}
				</button>
			</div>

			<button class=" bg-red-500 w-full mt-4 p-2 lg:text-lg font-medium text-white rounded-md"
				>Sign up</button
			>
			<p class=" text-center mt-4">
				Already have a account? <a href="/signin" class=" font-semibold text-red-500">Sign in</a>
			</p>
		</form>
		<p class=" text-center mt-4"><a href="/">&lt; Back</a></p>
	</div>
</div>
