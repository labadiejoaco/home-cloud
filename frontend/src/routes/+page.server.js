import { error, fail } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export async function load({ fetch }) {
	try {
		const res = await fetch('http://localhost:8080/api/');

		const data = await res.json();

		const { files, directories } = data.content;

		return {
			files,
			directories
		};
	} catch (err) {
		throw error(500, {
			message: 'Failed to fetch directories and files'
		});
	}
}

/** @type {import('./$types').Actions} */
export const actions = {
	create: async (event) => {
		const form = await event.request.formData();
		const name = form.get('name');

		try {
			const res = await fetch('http://localhost:8080/api/directory/', {
				method: 'POST',
				body: JSON.stringify({
					name
				}),
				headers: {
					'Content-type': 'application/json'
				}
			});

			const data = await res.json();

			if (data.error) {
				return fail(400, { form: 'create-directory', error: data.error });
			}

			return { success: true, form: 'create-directory', message: data.message };
		} catch (err) {
			return fail(500, { form: 'create-directory', error: 'Internal Server Error' });
		}
	},
	upload: async (event) => {
		const form = await event.request.formData();

		try {
			const res = await fetch('http://localhost:8080/api/upload/', {
				method: 'POST',
				body: form
			});

			const data = await res.json();

			if (data.error) {
				return fail(400, { form: 'upload-files', error: data.error });
			}

			return { success: true, form: 'upload-files', message: data.message };
		} catch (err) {
			return fail(500, { form: 'upload-files', error: 'Internal Server Error' });
		}
	}
};
