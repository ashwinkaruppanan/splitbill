<script>
	import { fly } from 'svelte/transition';
	import { goto } from '$app/navigation';
	import { api } from '../../api';
	let serverErrorMessage = null;

	let showPassword = false;
	let signupData = {
		name: '',
		email: '',
		password: ''
	};
	let showPasswordDescription = false;
	let error = {
		serverError: false,
		clientError: false,
		name: false,
		email: false,
		password: false
	};

	$: hasSixOrMoreChars = /^.{6,}$/.test(signupData.password);
	$: hasOneNumber = /.*\d.*/.test(signupData.password);
	$: hasOneLetter = /.*[a-zA-Z].*/.test(signupData.password);
	$: hasOneSpecialChar = /.*[^a-zA-Z0-9].*/.test(signupData.password);

	const handleChange = () => {
		error = {
			serverError: false,
			clientError: false,
			name: false,
			email: false,
			password: false
		};
	};

	const handleSubmit = async () => {
		if (signupData.name.length < 2 || !/^\w+( \w+)?$/.test(signupData.name)) {
			error.name = true;
			error.clientError = true;
		}
		if (!/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/.test(signupData.email)) {
			error.email = true;
			error.clientError = true;
		}
		if (!hasSixOrMoreChars || !hasOneNumber || !hasOneLetter || !hasOneSpecialChar) {
			error.password = true;
			error.clientError = true;
		}

		if (error.clientError) {
			return;
		}

		try {
			console.log(`${api}/api/v1/signup`);
			const response = await fetch(`${api}/api/v1/signup`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(signupData)
			});

			const data = await response.json();
			if (response.ok) {
				goto('/app');
			} else {
				serverErrorMessage = data.error;
				error.serverError = true;
			}
		} catch (error) {
			console.log(error);
		}
	};
</script>

<div class="flex justify-center md:mt-32 mt-12">
	<div class=" md:w-[28rem] lg:h-[25rem] mx-5">
		<h1 class=" sm:text-2xl text-xl font-semibold text-center">
			Split<span class=" text-red-500">Bill</span>
		</h1>
		<h1 class="lg:text-xl mt-2 text-center">Start with your free account today</h1>
		<form class=" mt-4" on:submit|preventDefault={handleSubmit}>
			<label for="name">Full Name</label>
			<input
				type="text"
				name="name"
				bind:value={signupData.name}
				on:input={(error.clientError || error.serverError) && handleChange}
				class="outline-none p-2 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
			/>
			{#if error.name}
				<p class="text-red-500">enter valid name</p>
			{/if}
			<div class=" mt-4">
				<label for="email"> Email address </label><br />
				<input
					type="text"
					name="email"
					bind:value={signupData.email}
					on:input={(error.clientError || error.serverError) && handleChange}
					class="outline-none p-2 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
				/>
				{#if error.email}
					<p class="text-red-500">enter valid email</p>
				{/if}
			</div>

			<div class=" relative mt-4">
				<label for="password">Password</label>
				{#if showPassword}
					<input
						type="text"
						name="password"
						bind:value={signupData.password}
						on:focus={() => (showPasswordDescription = true)}
						on:input={(error.clientError || error.serverError) && handleChange}
						class="outline-none p-2 pr-16 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
					/>
				{:else}
					<input
						type="password"
						name="password"
						bind:value={signupData.password}
						on:focus={() => (showPasswordDescription = true)}
						on:input={(error.clientError || error.serverError) && handleChange}
						class="outline-none p-2 pr-16 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
					/>
				{/if}
				<button
					type="button"
					class="absolute right-2 bottom-2"
					on:click={() => {
						showPassword = !showPassword;
					}}
				>
					{showPassword ? 'hide' : 'show'}
				</button>
			</div>
			{#if error.password}
				<p class="text-red-500">enter valid password</p>
			{/if}

			{#if showPasswordDescription}
				<div class="px-7 mt-4" transition:fly={{ y: '-10', delay: '100' }}>
					<ul class="list-disc">
						<li class={hasSixOrMoreChars ? 'text-green-600' : undefined}>6 or more characters</li>
						<li class={hasOneNumber ? 'text-green-600' : undefined}>One number</li>
						<li class={hasOneLetter ? 'text-green-600' : undefined}>One letter</li>
						<li class={hasOneSpecialChar ? 'text-green-600' : undefined}>One special character</li>
					</ul>
				</div>
			{/if}

			<button
				type="submit"
				class=" bg-red-500 w-full mt-4 p-2 lg:text-lg font-medium text-white rounded-md"
				>Sign up</button
			>
			<p class=" text-center mt-4">
				Already have a account? <a href="/signin" class=" font-semibold text-red-500">Sign in</a>
			</p>
		</form>
		<p class=" text-center mt-4"><a href="/">&lt; Back</a></p>
		{#if error.serverError}
			<p class=" text-center text-red-500 mt-6 text-xl">{serverErrorMessage}</p>
		{/if}
	</div>
</div>
