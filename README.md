# DetFes

Simple ReST API ([fiber](https://github.com/gofiber)) to check if image had human face or not (i'm using [go-face](https://github.com/Kagami/go-face)).

POST to <baseurl>/indentify with multipart/form name: `image` (you can change it via config.yaml)


But, first download required data in `models` folder:
```
$ cd models
$ wget https://github.com/Kagami/go-face-testdata/raw/master/models/shape_predictor_5_face_landmarks.dat
$ wget https://github.com/Kagami/go-face-testdata/raw/master/models/dlib_face_recognition_resnet_model_v1.dat
$ wget https://github.com/Kagami/go-face-testdata/raw/master/models/mmod_human_face_detector.dat
```
