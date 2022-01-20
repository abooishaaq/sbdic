<script>
    export let title;
    export let desc;
    export let count;
    export let tagsStr;
    export let btnText;
    export let group;
    export let formSubmit = () => {};

    export let groups = [{ title: "", id: -1 }];

    if (group && group.title) groups.push(group);

    let groupValue = group?.title;

    $: group = groups.filter((g) => g.title === groupValue)[0];

    const uniqGroups = (a) => {
        const seen = {};
        return a.filter((item) =>
            seen.hasOwnProperty(item.title) ? false : (seen[item.title] = true)
        );
    };

    fetch("/api/list-groups")
        .then((res) => {
            console.log(res);
            return res.json();
        })
        .then((res) => {
            if (res) {
                groups = [...groups, ...res];
                groups = uniqGroups(groups);
            }
        });
</script>

<form on:submit={formSubmit}>
    <label for="title">title</label>
    <input bind:value={title} name="title" placeholder="title" />
    <label for="description">description</label>
    <textarea
        bind:value={desc}
        name="description"
        placeholder="description"
        rows="10"
    />
    <label for="tags">tags</label>
    <input bind:value={tagsStr} placeholder="tags separated by space" />
    <label for="count">count</label>
    <input
        bind:value={count}
        placeholder="count"
        type="number"
        name="count"
        min="0"
    />
    <label for="group">group</label>
    <select name="group" bind:value={groupValue}>
        {#each groups as g}
            <option selected={group?.title === g.title}>{g.title}</option>
        {/each}
    </select>
    <button type="submit">{btnText}</button>
</form>

<style>
    form {
        display: grid;
        row-gap: 12px;
        grid-template-columns: 1fr 3fr;
    }

    textarea {
        min-width: 100%;
        max-width: 100%;
        min-height: 100px;
        max-height: 100px;
    }

    button {
        min-width: 320px;
        grid-column: span 2;
        justify-self: center;
    }
</style>
