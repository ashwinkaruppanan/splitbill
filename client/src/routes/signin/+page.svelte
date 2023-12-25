<script>
	import { api } from '../../api';
	import { goto } from '$app/navigation';

	let showPassword = false;
	let serverErrorMessage = null;

	let error = {
		serverError: false,
		clientError: false,
		name: false,
		email: false,
		password: false
	};

	let signinData = {
		email: '',
		password: ''
	};

	$: hasSixOrMoreChars = /^.{6,}$/.test(signinData.password);
	$: hasOneNumber = /.*\d.*/.test(signinData.password);
	$: hasOneLetter = /.*[a-zA-Z].*/.test(signinData.password);
	$: hasOneSpecialChar = /.*[^a-zA-Z0-9].*/.test(signinData.password);

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
		if (!/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/.test(signinData.email)) {
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
			const response = await fetch(`${api}/api/v1/signin`, {
				method: 'POST',
				body: JSON.stringify(signinData)
			});
			const data = await response.json();
			if (response.ok) {
				goto('/app');
			} else {
				serverErrorMessage = data.error;
				error.serverError = true;
			}
		} catch (e) {
			console.log(e);
		}
	};
</script>

<div class="flex justify-center md:mt-32 mt-12">
	<div class=" md:w-[28rem] lg:h-[25rem] mx-5">
		<h1 class=" sm:text-2xl text-xl font-semibold text-center">
			Split<span class=" text-red-500">Bill</span>
		</h1>
		<h1 class="lg:text-xl mt-2 text-center">Sign in to your account</h1>
		<form class=" mt-4" on:submit|preventDefault={handleSubmit}>
			<label for="email"> Email address </label><br />
			<input
				type="text"
				name="email"
				bind:value={signinData.email}
				on:input={(error.clientError || error.serverError) && handleChange}
				class="outline-none p-2 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
			/>
			{#if error.email}
				<p class="text-red-500">enter valid email</p>
			{/if}
			<br />
			<div class=" flex justify-between mt-4">
				<label for="password">Password</label>
				<a href="/password" class="  hover:text-red-500">Forgot password?</a>
			</div>
			<div class=" relative">
				{#if showPassword}
					<input
						type="text"
						name="password"
						bind:value={signinData.password}
						on:input={(error.clientError || error.serverError) && handleChange}
						class="outline-none p-2 pr-16 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
					/>
				{:else}
					<input
						type="password"
						name="password"
						bind:value={signinData.password}
						on:input={(error.clientError || error.serverError) && handleChange}
						class="outline-none p-2 pr-16 ring-1 ring-gray-500 w-full mt-1 focus:ring-red-400 rounded-md"
					/>
				{/if}

				<button
					type="button"
					class=" absolute right-2 bottom-2"
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

			<button
				type="submit"
				class=" bg-red-500 w-full mt-4 p-2 lg:text-lg font-medium text-white rounded-md"
				>Sign in</button
			>
			<p class=" text-center mt-4">
				Don't have a account? <a href="/signup" class=" font-semibold text-red-500">Sign up</a>
			</p>
		</form>
		<p class=" text-center mt-4"><a href="/">&lt; Back</a></p>
		{#if error.serverError}
			<p class=" text-center text-red-500 mt-6 text-xl">{serverErrorMessage}</p>
		{/if}
	</div>
</div>
