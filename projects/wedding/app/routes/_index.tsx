import type { ActionFunction, ActionFunctionArgs, LoaderFunction } from "@remix-run/cloudflare";
import { json } from "@remix-run/cloudflare";
import { Form, useLoaderData } from "@remix-run/react";

interface Env {
  MY_DB: D1Database;
}

type rsvp = {
    rsvp_id:number,
    first_name:number,
    last_name:number,
    plus_one:number,
    plus_one_first_name:number,
    plus_one_last_name:number,
    coming_to_thailand:number,
    coming_to_uk:number,
    diet:number
}

export const loader: LoaderFunction = async ({ context, params })  => {
  const env = context.cloudflare.env as Env
  const { results } = await env.MY_DB.prepare("SELECT * FROM rsvp ORDER BY rowid desc LIMIT 1").all();
  return json(results);
};

export const action: ActionFunction = async ({request, context} : ActionFunctionArgs) => {
    const env = context.cloudflare.env

    const formData = await request.formData()
    const    first_name = "first_name"
    const    last_name = "last_name"
    const    plus_one = 1
    const    plus_one_first_name = "plus_one_first_name"
    const    plus_one_last_name = "plus_one_last_name"
    const    coming_to_thailand = 2
    const    coming_to_uk = 2
    const    diet = "diet"

    console.log(formData)

    return await env.MY_DB.prepare(`
        INSERT INTO rsvp (first_name, last_name, plus_one, plus_one_first_name, plus_one_last_name, coming_to_thailand, coming_to_uk, diet)
        VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8);
    `).bind(
        first_name,
        last_name,
        plus_one,
        plus_one_first_name,
        plus_one_last_name,
        coming_to_thailand,
        coming_to_uk,
        diet,
    ).run().then((resp) => {
        return resp
    })

}

export default function Index() {
  const results: rsvp[] = useLoaderData<typeof loader>() as rsvp[];
  return (
    <div style={{ fontFamily: "system-ui, sans-serif", lineHeight: "1.8" }}>
      <h1>Welcome to Remix</h1>
      <div>
        A value from D1:
        <pre><code>{JSON.stringify(results, null, 3)}</code></pre>
      </div>

      <div>
        <Form method="POST">
        <div>
          <label>
            First Name: <input type="text" name="first-name" required />
          </label>
        </div>
        <div>
          <label>
            Last Name: <input type="text" name="last-name" required />
          </label>
        </div>
        <div>
          <label>
            Plus One: <input type="text" name="plus-one" required />
          </label>
        </div>
        <div>
          <label>
            Plus One First Name: <input type="text" name="plus-one-first-name" required />
          </label>
        </div>
        <div>
          <label>
            Plus One Last Name: <input type="text" name="plus-one-first-name" required />
          </label>
        </div>
        <div>
          <label>
            Coming To Thailand: <input type="text" name="coming-to-thailand" required />
          </label>
        </div>
        <div>
          <label>
            Coming To UK: <input type="text" name="coming-to-uk" required />
          </label>
        </div>
        <div>
          <label>
            Diet restrictions: <input type="text" name="diet" required />
          </label>
        </div>
        <button type="submit">Add Resource</button>
      </Form>
      </div>
    </div>
  );
}
