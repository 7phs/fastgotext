#include <iostream>
#include "../fasttext/src/fasttext.h"

const int RES_OK = 0;
const int RES_ERROR_NOT_OPEN = 1;
const int RES_ERROR_WRONG_MODEL = 2;
const int RES_ERROR_NOT_INIT = 3;

extern "C" {
    struct WrapperVector {
        fasttext::Vector *vector;
    };

    struct WrapperDictionary {
        std::shared_ptr<const fasttext::Dictionary> dict;
    };

    struct WrapperFastText {
        fasttext::FastText *model;
    };
}

bool checkModelInitialization(const struct WrapperFastText* wrapper) {
    if (wrapper==nullptr ||
        wrapper->model==nullptr ||
        wrapper->model->getDictionary()==nullptr
    ) {
        return false;
    }

    return true;
}

bool checkModelFile(std::istream& in) {
    int32_t magic, version;

    in.read((char*)&(magic), sizeof(int32_t));
    if (magic != FASTTEXT_FILEFORMAT_MAGIC_INT32) {
        return false;
    }

    in.read((char*)&(version), sizeof(int32_t));
    if (version > FASTTEXT_VERSION) {
        return false;
    }

    return true;
}

const int checkVectorsFile(const std::string& path, const int ndim) {
    std::ifstream in(path);

    int64_t n, dim;
    if (!in.is_open()) {
        return RES_ERROR_NOT_OPEN;
    }

    in >> n >> dim;
    if (dim != ndim) {
        return RES_ERROR_WRONG_MODEL;
    }

    in.close();

    return RES_OK;
}

struct WrapperVector* Vector(int ndim) {
    WrapperVector *wrapper = (WrapperVector *)malloc(sizeof (struct WrapperVector));

    wrapper->vector = new fasttext::Vector(ndim);

    return wrapper;
}

struct WrapperDictionary* Dictionary(std::shared_ptr<const fasttext::Dictionary> dict) {
    WrapperDictionary *wrapper = (WrapperDictionary *)malloc(sizeof (struct WrapperDictionary));

    wrapper->dict = dict;

    return wrapper;
}

struct WrapperVector* Subtract(struct WrapperVector* vec1, struct WrapperVector* vec2) {
    WrapperVector *result = Vector(vec1->vector->m_);

    if (vec1->vector->m_!=vec2->vector->m_) {
        return result;
    }

    for(unsigned int i=0, sz = vec1->vector->m_; i<sz; i++) {
        result->vector->data_[i] = vec1->vector->data_[i] - vec2->vector->data_[i];
    }

    return result;
}

 struct WrapperVector* Pow(struct WrapperVector* wrapper) {
    for(unsigned int i=0, sz = wrapper->vector->m_; i<sz; i++) {
        wrapper->vector->data_[i] *= wrapper->vector->data_[i];
    }

    return wrapper;
}

float Sum(struct WrapperVector* wrapper) {
    float result = 0.;


    for(unsigned int i=0, sz = wrapper->vector->m_; i<sz; i++) {
        result += wrapper->vector->data_[i];
    }

    return result;
}

extern "C" {
    int DICT_Find(struct WrapperDictionary* wrapper, const char* word) {
        return int(wrapper->dict->getId(word));
    }

    void DICT_Release(struct WrapperDictionary* wrapper) {
        wrapper->dict.reset();

        free(wrapper);
    }

    void VEC_Release(struct WrapperVector* wrapper) {
        delete wrapper->vector;

        free(wrapper);
    }

    float VEC_Distance(struct WrapperVector* vec1, struct WrapperVector* vec2) {
        struct WrapperVector* vec = Pow(Subtract(vec1, vec2));

        float result = Sum(vec);

        VEC_Release(vec);

        return result;
    }

    struct WrapperFastText* FastText() {
        WrapperFastText *wrapper = (WrapperFastText *)malloc(sizeof (struct WrapperFastText));

        wrapper->model = new fasttext::FastText();

        return wrapper;
    }

    const int FT_LoadModel(struct WrapperFastText* wrapper, const char* path) {
        std::ifstream ifs(path, std::ifstream::binary);

        if (!ifs.is_open()) {
            return RES_ERROR_NOT_OPEN;
        }

        if (!checkModelFile(ifs)) {
            return RES_ERROR_WRONG_MODEL;
        }

        wrapper->model->loadModel(ifs);

        ifs.close();

        return RES_OK;
    }

    const int FT_LoadVectors(struct WrapperFastText* wrapper, const char* path) {
        std::string vectorsPath(path);

        if (!checkModelInitialization(wrapper)) {
            return RES_ERROR_NOT_INIT;
        }

        const int res = checkVectorsFile(vectorsPath, wrapper->model->getDimension());
        if (res!=RES_OK) {
            return res;
        }

        wrapper->model->loadVectors(vectorsPath);

        return RES_OK;
    }

    struct WrapperDictionary* FT_GetDictionary(struct WrapperFastText* wrapper) {
        return Dictionary(wrapper->model->getDictionary());
    }

    struct WrapperVector* FT_GetVector(struct WrapperFastText* wrapper, const char* word) {
        struct WrapperVector* wrap_vector = Vector(wrapper->model->getDimension());

        wrapper->model->getVector(*wrap_vector->vector, word);

        return wrap_vector;
    }

    void FT_Release(struct WrapperFastText* wrapper) {
        delete wrapper->model;

        free(wrapper);
    }
}