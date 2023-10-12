<script>
    import { env } from '$env/dynamic/public';

    let counter = 1;
    let item = 'address';
    let result = null;

    let apiBaseUrl = env.PUBLIC_API_BASE_URL

    async function executeRequest() {
        try {
            const response = await fetch(`${apiBaseUrl}/random/${item}?count=${counter}`);
            const data = await response.json();
            if (data.status === 'ok') {
                result = data.data
            } else {
                result = [ data.message ];
            }
        } catch (error) {
            result = [ 'Error: ' + error.message ];
        }
    }
</script>

<main>
    <h1 title="Backend on {apiBaseUrl}">Random Generator</h1>
    <div>
        <label for="itemDropdown">Type</label>
        <select id="itemDropdown" bind:value={item}>
            <option value="adjective">Adjective</option>
            <option value="address">Address</option>
            <option value="city">City</option>
            <option value="country">Country</option>
            <option value="currency">Currency</option>
            <option value="day">Day</option>
            <option value="email">Email</option>
            <option value="firstnamem">First Name (Male)</option>
            <option value="firstnamef">First Name (Female)</option>
            <option value="fulldate">Full date</option>
            <option value="fullnamem">Full Name (Male)</option>
            <option value="fullnamef">Full Name (Female)</option>
            <option value="ipv4address">Ipv4address</option>
            <option value="ipv6address">Ipv6address</option>
            <option value="lastname">Lastname</option>
            <option value="locale">Locale</option>
            <option value="month">Month</option>
            <option value="paragraph">Paragraph</option>
            <option value="phonenumber">Phone number</option>
            <option value="sillyname">Silly name</option>
            <option value="state">State</option>
            <option value="street">Street</option>
            <option value="useragent">User Agent</option>
        </select>
    </div>
    <div>
        <label for="counterDropdown">Counter</label>
        <select id="counterDropdown" bind:value={counter}>
            {#each Array.from({ length: 10 }, (_, i) => i + 1) as num}
                <option value={num}>{num}</option>
            {/each}
        </select>
    </div>
    <div>
        <button on:click={executeRequest}>Generate</button>
    </div>
    {#if result !== null}
        <div class="result-container">
            {#each result as item}
                <div class="result">{item}</div>
            {/each}
        </div>
    {/if}
</main>

<style>
    main {
        text-align: center;
        padding: 1em;
        max-width: 640px;
        margin: 0 auto;
    }

    h1 {
        color: #ff3e00;
        text-transform: uppercase;
        font-size: 4em;
        font-weight: 100;
    }

    label {
        margin-bottom: 10px;
        font-weight: bold;
    }

    div {
        margin-bottom: 15px;
    }

    div.result-container {
        margin: 40px;
        padding: 10px;
        border: 1px solid #333333;
    }

    div.result {
        padding: 10px;
        margin: 0;
    }

    select {
        width: 200px;
        text-align: center;
    }

    button {
        width: 200px;
        text-align: center;
    }
</style>
