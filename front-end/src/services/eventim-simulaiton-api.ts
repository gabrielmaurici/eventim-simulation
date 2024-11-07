import axios, { AxiosError } from 'axios';

const api = axios.create({
  baseURL: 'http://localhost/',
  headers: {
    'Content-Type': 'application/json',
  },
});

const eventimSimulationMethodsApi = {
  async post<I, R>(endpoint: string, data?: I | null): Promise<R | null> {
    try {
      const response = await api.post(endpoint, data);
      return response.data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const status = error.response?.status;
        if (status === 500 || status == 400) {
          throw new Error(error.response?.data);
        }
      }

      throw error;
    }
  },
};

export default eventimSimulationMethodsApi;