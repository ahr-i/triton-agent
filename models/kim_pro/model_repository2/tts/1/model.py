from transformers import VitsModel, AutoTokenizer
import torch
import triton_python_backend_utils as pb_utils
#from scipy.io import wavfile
import json
import numpy
import time
import io
import base64
from scipy.io.wavfile import write

class TritonPythonModel:
    def initialize(self, args):
        self.model = VitsModel.from_pretrained("facebook/mms-tts-eng")
        self.tokenizer = AutoTokenizer.from_pretrained("facebook/mms-tts-eng")
    def execute(self, requests):
        responses = []
        for request in requests:
            text_data = pb_utils.get_input_tensor_by_name(request, "text").as_numpy()
            text = text_data[0].decode('utf-8')
            inputs = self.tokenizer(text, return_tensors="pt")
            with torch.no_grad():
                output = self.model(**inputs).waveform
            sampling_rate = self.model.config.sampling_rate
            print(sampling_rate)

            inference_response = pb_utils.InferenceResponse(
                output_tensors=[
                    pb_utils.Tensor(
                        "audio_res",
                        numpy.array([output.float()]),
                    )       
                ]
            )
            responses.append(inference_response)
            
        return responses


        def save_response_audio(self, output):
            unique_id = int(time.time() * 1000)
            save_dir = "./files/"
            file_name = f"audio_res_{unique_id}.wav"
            audio_data_int16 = (output.float().numpy() * 32767).astype(numpy.int16)
            scipy.io.wavfile.write(file_name, rate=self.model.config.sampling_rate, data=audio_data_int16)
            print(file_name)
