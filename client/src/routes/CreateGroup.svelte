<script>
    let title = "";
    let desc = "";

    const formSubmit = (event) => {
        event.preventDefault();
        if (!title) return alert("title cannot be blank");
        fetch("/api/create-group", {
            method: "POST",
            body: JSON.stringify({ title, desc }),
        })
            .then((res) => {
                console.log(res);
                if (res.status === 200) {
                    alert("group created successfully");
                    title = "";
                    desc = "";
                    count = 1;
                    tags = [];
                    tagsStr = "";
                } else return res.json();
            })
            .then((res) => alert(res.err));
    };
</script>

<div id="container">
    <h1>Create Group</h1>

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
        <button type="submit">Create Group</button>
    </form>
</div>

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
